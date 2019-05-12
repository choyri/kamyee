package kyrabbitmq

import "github.com/streadway/amqp"

// 生产者产物
type Message struct {
	Header amqp.Table
	Body   interface{}
}

// 消费者产物
type Publication struct {
	RoutingKey string
	Body       []byte
}

// 消费者处理者
type ConsumerHandler struct {
	QueueName string
	BindKey   string
	FuncLists []func(*Publication) error
}
