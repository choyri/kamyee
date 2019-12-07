package rabbitmq

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type channel struct {
	Channel *amqp.Channel
	tag     string // Consumer tag
}

func newChannel(conn *connection) (*channel, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	ret := channel{
		tag: id.String(),
	}

	ret.Channel, err = conn.Connection.Channel()
	if err != nil {
		return nil, err
	}

	if conn.prefetchCount > 0 {
		err = ret.Channel.Qos(int(conn.prefetchCount), 0, true)
		if err != nil {
			return nil, err
		}
	}

	return &ret, nil
}

func (ch *channel) Close() error {
	if ch.Channel == nil {
		return nullChannel
	}

	return ch.Channel.Close()
}

func (ch *channel) Publish(exchange, routingKey string, message amqp.Publishing) error {
	return ch.Channel.Publish(exchange, routingKey, false, false, message)
}

func (ch *channel) Consume(queue string) (<-chan amqp.Delivery, error) {
	return ch.Channel.Consume(
		queue,  // Name
		ch.tag, // consumerTag,
		false,  // noAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // arguments
	)
}

func (ch *channel) DeclareExchange(exchange exchange) error {
	return ch.Channel.ExchangeDeclare(
		exchange.Name, // name
		exchange.Type, // kind
		true,          // durable
		false,         // autoDelete
		false,         // internal
		false,         // noWait
		nil,           // args
	)
}

func (ch *channel) DeclareQueue(name string) error {
	_, err := ch.Channel.QueueDeclare(
		name,  // Name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)

	return err
}

func (ch *channel) BindQueue(queue, routingKey, exchange string) error {
	return ch.Channel.QueueBind(
		queue,      // Name of the queue
		routingKey, // routingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	)
}
