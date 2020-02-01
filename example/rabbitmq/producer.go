package main

import (
	"fmt"
	"github.com/choyri/kamyee/example"
	"github.com/choyri/kamyee/rabbitmq"
	"time"
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

	err = p.PublishBody("custom.fine", "hello, my friend, I'm very ok.")
	example.HandleErr(err)

	pDelayed, err := rabbitmq.NewProducer(rabbitmq.DelayedExchange("test_delayed"))
	example.HandleErr(err)

	fmt.Printf("delayed message begin publish: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	err = pDelayed.PublishBody("delayed.ok", "hello, my friend, I'm ok, from delayed message, after 6s.", 6*time.Second)
	example.HandleErr(err)
}
