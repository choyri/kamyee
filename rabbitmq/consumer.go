package rabbitmq

import (
	"fmt"
	"sync"
	"time"
)

type Consumer struct {
	r        *rabbitmq
	handlers []ConsumerHandler
	channels []*channel
	shutdown bool

	sync.Mutex
}

func NewConsumer(opts ...Option) (*Consumer, error) {
	r := newRabbitMQ(opts...)

	if err := r.Connect(); err != nil {
		return nil, fmt.Errorf("new producer failed: %w", err)
	}

	ret := Consumer{
		r: r,
	}

	return &ret, nil
}

func (c *Consumer) RegisterHandler(handlers []ConsumerHandler) error {
	queueMap := make(map[string]int)

	for _, v := range handlers {
		if len(v.Queue) == 0 {
			return fmt.Errorf("empty queue (routingKey: %s)", v.RoutingKey)
		}

		queueMap[v.Queue]++

		if queueMap[v.Queue] > 1 {
			return fmt.Errorf("duplicate queue (queue: %s)", v.Queue)
		}
	}

	c.handlers = handlers

	return nil
}

func (c *Consumer) Startup() error {
	if c.r.Conn == nil {
		return fmt.Errorf("startup failed: %w", nullConnection)
	}

	var err error

	for _, v := range c.handlers {
		err = c.r.Conn.DeclareAndBindQueue(v.Queue, v.RoutingKey)
		if err != nil {
			return err
		}

		go c.handle(v.Queue, v.Handler)
	}

	return nil
}

func (*Consumer) Keepalive() {
	select {}
}

func (c *Consumer) Shutdown() error {
	c.Lock()
	defer c.Unlock()

	c.shutdown = true

	for _, v := range c.channels {
		err := v.Channel.Cancel(v.tag, true)
		if err != nil {
			return fmt.Errorf("consumer shutdown failed: %w", err)
		}
	}

	if err := c.r.Close(); err != nil {
		return fmt.Errorf("consumer shutdown failed: %w", err)
	}

	return nil
}

func (c *Consumer) handle(queue string, handler func(*Publication) error) {
	var (
		minDelay  = time.Second
		maxDelay  = 30 * time.Second
		expFactor = time.Duration(2)
		delay     = minDelay
	)

	for {
		// 检查消费者是否已发起停止操作
		c.Lock()
		shutdown := c.shutdown
		c.Unlock()
		if shutdown {
			return
		}

		select {
		case <-c.r.Conn.closeChan:
			// 连接已关闭 退出消费
			return
		case <-c.r.Conn.waitChan:
			// 重连时会在此阻塞
		}

		c.r.Lock()

		if c.r.Conn.connected == false {
			c.r.Unlock()
			continue
		}

		channel, deliveries, err := c.r.Conn.Consume(queue)
		c.channels = append(c.channels, channel)

		c.r.Unlock()

		switch err {
		case nil:
			delay = minDelay
		default:
			// 如果发生了错误会进到这个分支 等待重试
			if delay > maxDelay {
				delay = maxDelay
			}

			time.Sleep(delay)

			delay *= expFactor
			continue
		}

		for delivery := range deliveries {
			p := Publication{
				RoutingKey: delivery.RoutingKey,
				Body:       delivery.Body,
			}

			if err := handler(&p); err == nil {
				_ = delivery.Ack(false)
			} else {
				_ = delivery.Nack(false, true)
			}
		}
	}
}
