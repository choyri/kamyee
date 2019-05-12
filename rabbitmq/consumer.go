package kyrabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Consumer struct {
	broker   *broker
	channel  *channel
	shutdown bool
	handlers []ConsumerHandler
	sync.Mutex
}

func NewConsumer(config *Config) (*Consumer, error) {
	broker := newBroker(config)

	if err := broker.Connect(); err != nil {
		return nil, WrapErr("New consumer failed", err)
	}

	return &Consumer{
		broker: broker,
	}, nil
}

func (entity *Consumer) Startup() error {
	if entity.broker.conn == nil {
		return WrapErr("Startup failed", errors.New("not connected"))
	}

	for _, handler := range entity.handlers {
		entity.broker.wg.Add(1)
		go entity.handle(&handler)
	}

	return nil
}

func (entity *Consumer) Shutdown() error {
	entity.Lock()
	defer entity.Unlock()

	entity.shutdown = true

	if entity.channel != nil {
		if err := entity.channel.Close(); err != nil {
			return WrapErr("consumer shutdown failed", err)
		}
	}

	return nil
}

func (entity *Consumer) RegisterHandler(handlers []ConsumerHandler) *Consumer {
	entity.handlers = handlers
	return entity
}

func (entity *Consumer) handle(handler *ConsumerHandler) {
	minDelay := 100 * time.Millisecond
	maxDelay := 30 * time.Second
	expFactor := time.Duration(2)
	delay := minDelay

	for {
		// 检查消费者是否已发起停止操作
		entity.Lock()
		shutdown := entity.shutdown
		entity.Unlock()
		if shutdown {
			return
		}

		select {
		case <-entity.broker.conn.closeChan:
			// 连接已关闭 退出消费
			return
		case <-entity.broker.conn.waitChan:
			// 重连时会在此阻塞
		}

		// 检查是否已连接
		entity.broker.Lock()
		if entity.broker.conn.connected == false {
			entity.broker.Unlock()
			continue
		}

		channel, deliveries, err := entity.broker.conn.Consume(handler.QueueName, handler.BindKey)
		entity.broker.Unlock()

		switch err {
		case nil:
			delay = minDelay
			entity.Lock()
			entity.channel = channel
			entity.Unlock()
		default:
			// 如果发生了错误会进到这个分支 等待重试
			if delay > maxDelay {
				delay = maxDelay
			}

			time.Sleep(delay)

			delay *= expFactor
			continue
		}

		// 开始处理事件
		var subWg sync.WaitGroup

		for delivery := range deliveries {
			subWg.Add(1)

			go func(d amqp.Delivery) {
				p := Publication{
					RoutingKey: d.RoutingKey,
					Body:       d.Body,
				}

				for _, fn := range handler.FuncLists {
					if err := fn(&p); err == nil {
						_ = d.Ack(false)
					} else {
						_ = d.Nack(false, true)
					}
				}

				subWg.Done()
			}(delivery)
		}

		subWg.Wait()

		entity.broker.wg.Done()
	}
}
