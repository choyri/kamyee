package rabbitmq

import (
	"sync"
)

type rabbitmq struct {
	Conn *connection
	opts Options

	sync.Mutex
}

func newRabbitMQ(opts ...Option) *rabbitmq {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if len(options.URL) == 0 {
		options.URL = DefaultAMQPURL
	}

	if len(options.Exchange.Name) == 0 {
		options.Exchange = DefaultExchange
	}

	return &rabbitmq{
		opts: options,
	}
}

func (r *rabbitmq) Connect() error {
	if r.Conn == nil {
		r.Conn = newConnection(r.opts)
	}

	return r.Conn.Connect()
}

func (r *rabbitmq) Close() error {
	if r.Conn == nil {
		return nullConnection
	}

	return r.Conn.Close()
}
