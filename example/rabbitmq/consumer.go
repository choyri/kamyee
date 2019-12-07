package main

import (
	"errors"
	"fmt"
	"github.com/choyri/kamyee/example"
	"github.com/choyri/kamyee/rabbitmq"
	"math/rand"
	"time"
)

func main() {
	c, err := rabbitmq.NewConsumer()
	example.HandleErr(err)

	err = c.RegisterHandler([]rabbitmq.ConsumerHandler{
		{
			Queue:      "test1",
			RoutingKey: "default.*",
			Handler: func(p *rabbitmq.Publication) error {
				fmt.Print("current # queue: test1, routingKey: default.*\n")
				fmt.Printf("receive data # routingKey: %s, get: %s\n\n", p.RoutingKey, string(p.Body))
				return nil
			},
		},
		{
			Queue:      "test2",
			RoutingKey: "default.fine",
			Handler: func(p *rabbitmq.Publication) error {
				fmt.Print("current # queue: test2, routingKey: default.fine\n")
				fmt.Printf("receive data # routingKey: %s, get: %s\n\n", p.RoutingKey, string(p.Body))
				return nil
			},
		},
		{
			Queue:      "test3",
			RoutingKey: "custom.*",
			Handler: func(p *rabbitmq.Publication) error {
				rand.Seed(time.Now().UnixNano())
				if rand.Intn(4) >= 3 {
					fmt.Print("current # queue: test3, routingKey: custom.*\n")
					fmt.Printf("receive data # routingKey: %s, get: %s\n\n", p.RoutingKey, string(p.Body))
					return nil
				}
				fmt.Printf("try requeue\n")
				return errors.New("err")
			},
		},
	})
	example.HandleErr(err)

	err = c.Startup()
	example.HandleErr(err)

	c.Keepalive()
}
