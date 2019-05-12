package kyrabbitmq

import (
	"errors"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type channel struct {
	uuid    string
	channel *amqp.Channel
}

func newChannel(conn *connection) (*channel, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	ret := channel{
		uuid: id.String(),
	}

	ret.channel, err = conn.connection.Channel()
	if err != nil {
		return nil, err
	}

	if conn.brokerConfig.PrefetchCount != 0 {
		if err = ret.channel.Qos(conn.brokerConfig.PrefetchCount, 0, conn.brokerConfig.PrefetchGlobal); err != nil {
			return nil, err
		}
	}

	return &ret, nil
}

func (entity *channel) Close() error {
	if entity.channel == nil {
		return errors.New("channel is nil")
	}

	return entity.channel.Close()
}

func (entity *channel) Publish(exchange, routingKey string, message *amqp.Publishing) error {
	if entity.channel == nil {
		return errors.New("channel is nil")
	}

	return entity.channel.Publish(exchange, routingKey, false, false, *message)
}

func (entity *channel) Consume(queue string) (<-chan amqp.Delivery, error) {
	return entity.channel.Consume(
		queue,       // name
		entity.uuid, // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
}

func (entity *channel) DeclareExchange(exchange *Exchange) error {
	return entity.channel.ExchangeDeclare(
		exchange.Name,  // name
		exchange.Type,  // kind
		true,           // durable
		false,          // autoDelete
		false,          // internal
		false,          // noWait
		*exchange.Args, // args
	)
}

func (entity *channel) DeclareQueue(name string) error {
	_, err := entity.channel.QueueDeclare(
		name,  // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)

	return err
}

func (entity *channel) BindQueue(queue, key, exchange string) error {
	return entity.channel.QueueBind(
		queue,    // name of the queue
		key,      // bindingKey
		exchange, // sourceExchange
		false,    // noWait
		nil,      // arguments
	)
}
