package rabbitmq

import (
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type connection struct {
	Connection *amqp.Connection
	Channel    *channel

	url           string
	exchange      exchange
	prefetchCount uint16
	closeChan     chan struct{}
	waitChan      chan struct{}

	connected bool
	sync.Mutex
}

type exchange struct {
	Name string
	Type string
	Args amqp.Table
}

func newConnection(opts Options) *connection {
	ret := connection{
		url:           opts.URL,
		exchange:      opts.Exchange,
		prefetchCount: opts.PrefetchCount,
		closeChan:     make(chan struct{}),
		waitChan:      make(chan struct{}),
	}

	close(ret.waitChan)

	return &ret
}

func (conn *connection) Connect() error {
	conn.Lock()

	if conn.connected {
		conn.Unlock()
		return nil
	}

	select {
	case <-conn.closeChan:
		// closeChan 被关闭时会进入该分支 故重建通道
		conn.closeChan = make(chan struct{})
	default:
		// 万事胜意
	}

	conn.Unlock()

	return conn.connect()
}

func (conn *connection) Close() error {
	conn.Lock()
	defer conn.Unlock()

	select {
	case <-conn.closeChan:
		return nil
	default:
		close(conn.closeChan)
		conn.connected = false
	}

	err := conn.Channel.Close()
	if err != nil {
		return err
	}

	return conn.Connection.Close()
}

func (conn *connection) Publish(routingKey string, msg amqp.Publishing) error {
	if conn.Channel == nil {
		return nullChannel
	}

	return conn.Channel.Publish(conn.exchange.Name, routingKey, msg)
}

func (conn *connection) Consume(queue string) (*channel, <-chan amqp.Delivery, error) {
	consumerChannel, err := newChannel(conn)
	if err != nil {
		return nil, nil, err
	}

	deliveries, err := consumerChannel.Consume(queue)
	if err != nil {
		return nil, nil, err
	}

	return consumerChannel, deliveries, nil
}

func (conn *connection) DeclareAndBindQueue(queue, routingKey string) error {
	if conn.Channel == nil {
		return nullChannel
	}

	err := conn.Channel.DeclareQueue(queue)
	if err != nil {
		return err
	}

	return conn.Channel.BindQueue(queue, routingKey, conn.exchange.Name)
}

func (conn *connection) connect() error {
	err := conn.tryConnect()
	if err != nil {
		return err
	}

	conn.Lock()
	conn.connected = true
	conn.Unlock()

	go conn.keepConnect()

	return nil
}

func (conn *connection) tryConnect() error {
	var err error

	conn.Connection, err = amqp.Dial(conn.url)
	if err != nil {
		return err
	}

	conn.Channel, err = newChannel(conn)
	if err != nil {
		return err
	}

	return conn.Channel.DeclareExchange(conn.exchange)
}

func (conn *connection) keepConnect() {
	connected := false

	for {
		if connected {
			if err := conn.tryConnect(); err != nil {
				time.Sleep(time.Second)
				continue
			}

			conn.Lock()
			conn.connected = true
			conn.Unlock()

			close(conn.waitChan)
		}

		connected = true

		notifyClose := make(chan *amqp.Error)
		conn.Connection.NotifyClose(notifyClose)

		select {
		case <-notifyClose:
			conn.Lock()
			conn.connected = false
			conn.waitChan = make(chan struct{})
			conn.Unlock()
		case <-conn.closeChan:
			return
		}
	}
}
