package kyrabbitmq

import (
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type connection struct {
	connection *amqp.Connection
	channel    *channel

	exchange     *Exchange
	brokerConfig *Config
	closeChan    chan struct{}
	waitChan     chan struct{}

	connected bool
	sync.Mutex
}

type Exchange struct {
	Name string
	Type string
	Args *amqp.Table
}

func newConnection(brokerConfig *Config) *connection {
	exchange := brokerConfig.Exchange

	if exchange == nil || exchange.Name == "" {
		exchange = &DefaultExchange
	} else if exchange.Type == "" {
		exchange.Type = DefaultExchange.Type
	}

	ret := connection{
		exchange:     exchange,
		brokerConfig: brokerConfig,
		closeChan:    make(chan struct{}),
		waitChan:     make(chan struct{}),
	}

	// 新建一个连接时 该通道要置关闭状态
	close(ret.waitChan)

	return &ret
}

func (entity *connection) Connect() error {
	entity.Lock()

	if entity.connected {
		entity.Unlock()
		return nil
	}

	select {
	case <-entity.closeChan:
		// 该通道被关闭时 会进入该分支 故重建通道
		entity.closeChan = make(chan struct{})
	default:
	}

	entity.Unlock()

	return entity.connect()
}

func (entity *connection) Close() error {
	entity.Lock()
	defer entity.Unlock()

	select {
	case <-entity.closeChan:
		return nil
	default:
		close(entity.closeChan)
		entity.connected = false
	}

	return entity.connection.Close()
}

func (entity *connection) Publish(routingKey string, msg *amqp.Publishing) error {
	return entity.channel.Publish(entity.exchange.Name, routingKey, msg)
}

func (entity *connection) Consume(queue, key string) (*channel, <-chan amqp.Delivery, error) {
	consumerChannel, err := newChannel(entity)
	if err != nil {
		return nil, nil, err
	}

	if err := consumerChannel.DeclareQueue(queue); err != nil {
		return nil, nil, err
	}

	if err := consumerChannel.BindQueue(queue, key, entity.exchange.Name); err != nil {
		return nil, nil, err
	}

	deliveries, err := consumerChannel.Consume(queue)
	if err != nil {
		return nil, nil, err
	}

	return consumerChannel, deliveries, nil
}

func (entity *connection) connect() error {
	if err := entity.tryConnect(); err != nil {
		return err
	}

	entity.Lock()
	entity.connected = true
	entity.Unlock()

	go entity.keepConnect()

	return nil
}

func (entity *connection) tryConnect() error {
	var err error

	if entity.connection, err = amqp.Dial(entity.brokerConfig.GetAmqpUrl()); err != nil {
		return err
	}

	if entity.channel, err = newChannel(entity); err != nil {
		return err
	}

	return entity.channel.DeclareExchange(entity.exchange)
}

func (entity *connection) keepConnect() {
	offline := false

	for {
		if offline {
			if err := entity.tryConnect(); err != nil {
				time.Sleep(time.Second)
				continue
			}

			entity.Lock()
			entity.connected = true
			entity.Unlock()

			close(entity.waitChan)
		}

		notifyClose := make(chan *amqp.Error)
		entity.connection.NotifyClose(notifyClose)

		select {
		case <-notifyClose:
			offline = true
			entity.Lock()
			entity.connected = false
			entity.waitChan = make(chan struct{})
			entity.Unlock()
		case <-entity.closeChan:
			return
		}
	}
}
