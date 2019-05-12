package kyrabbitmq

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
)

type Producer struct {
	broker *broker
}

func NewProducer(config *Config) (*Producer, error) {
	broker := newBroker(config)

	if err := broker.Connect(); err != nil {
		return nil, WrapErr("New producer failed", err)
	}

	return &Producer{
		broker: broker,
	}, nil
}

func (entity *Producer) Publish(routingKey string, message *Message) error {
	if entity.broker.conn == nil {
		return WrapErr("Publish failed", errors.New("not connected"))
	}

	body, err := json.Marshal(message.Body)
	if err != nil {
		return WrapErr("Marshal message failed", err)
	}

	msg := amqp.Publishing{
		Headers:     message.Header,
		ContentType: "application/json",
		Body:        body,
	}

	if err := entity.broker.conn.Publish(routingKey, &msg); err != nil {
		return WrapErr("Publish failed", err)
	}

	return nil
}
