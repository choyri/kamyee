package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
)

var (
	DefaultAMQPURL  = "amqp://guest:guest@127.0.0.1:5672/"
	DefaultExchange = exchange{
		Name: "kamyee",
		Type: amqp.ExchangeTopic,
	}

	nullChannel    = errors.New("channel is null")
	nullConnection = errors.New("connection is null")
)
