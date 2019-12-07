package main

import (
	"github.com/choyri/kamyee/example"
	"github.com/choyri/kamyee/rabbitmq"
)

func main() {
	p, err := rabbitmq.NewProducer()
	example.HandleErr(err)

	err = p.Publish("default.ok", rabbitmq.Message{
		Body: "hello, my friend, I'm ok.",
	})
	example.HandleErr(err)

	err = p.Publish("default.fine", rabbitmq.Message{
		Body: "hello, my friend, I'm fine.",
	})
	example.HandleErr(err)

	err = p.Publish("custom.ok", rabbitmq.Message{
		Body: "hello, my friend, I'm very ok.",
	})
	example.HandleErr(err)
}
