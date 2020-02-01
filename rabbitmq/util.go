package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
)

var (
	DefaultAMQPURL  = "amqp://guest:guest@127.0.0.1:5672/"
	DefaultExchange = exchange{
		Name: "kamyee",
		Kind: amqp.ExchangeTopic,
	}

	delayHeader            = "x-delay"
	delayedExchangeType    = "x-delayed-message"
	delayedExchangeArgsKey = "x-delayed-type"

	nullChannel    = errors.New("channel is null")
	nullConnection = errors.New("connection is null")
)
