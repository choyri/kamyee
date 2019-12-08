package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"time"
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

func (p *Producer) PublishBody(routingKey string, body interface{}, delayTime ...time.Duration) error {
	msg := Message{Body: body}

	if len(delayTime) > 0 {
		if p.r.Conn.exchange.Type != delayedExchangeType {
			return errors.New("cannot publish delayed message on a non-delay exchange")
		}
		msg.Header = amqp.Table{
			delayHeader: int(delayTime[0] / time.Millisecond),
		}
	}

	return p.Publish(routingKey, msg)
}
