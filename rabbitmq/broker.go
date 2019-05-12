package kyrabbitmq

import (
	"errors"
	"sync"
)

type broker struct {
	conn   *connection
	config *Config

	// 用于等待消费者处理事件
	wg sync.WaitGroup

	sync.Mutex
}

func newBroker(config *Config) *broker {
	return &broker{
		config: config,
	}
}

func (entity *broker) Connect() error {
	if entity.conn == nil {
		entity.conn = newConnection(entity.config)
	}

	return entity.conn.Connect()
}

func (entity *broker) Close() error {
	if entity.conn == nil {
		return errors.New("channel is nil")
	}

	err := entity.conn.Close()
	entity.wg.Wait()

	if err != nil {
		return err
	}

	return nil
}
