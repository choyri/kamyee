package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Producer struct {
	r *rabbitmq
}

func NewProducer(opts ...Option) (*Producer, error) {
	r := newRabbitMQ(opts...)

	if err := r.Connect(); err != nil {
		return nil, fmt.Errorf("new producer failed: %w", err)
	}

	ret := Producer{
		r: r,
	}

	return &ret, nil
}

func (p *Producer) Publish(routingKey string, message Message) error {
	if p.r.Conn == nil {
		return fmt.Errorf("publish failed: %w", nullConnection)
	}

	body, err := json.Marshal(message.Body)
	if err != nil {
		return fmt.Errorf("marshal message failed: %w", err)
	}

	msg := amqp.Publishing{
		Headers:     message.Header,
		ContentType: "application/json",
		Body:        body,
	}

	err = p.r.Conn.Publish(routingKey, msg)
	if err != nil {
		return fmt.Errorf("publish failed: %w", err)
	}

	return nil
}
