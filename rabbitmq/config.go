package kyrabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

var (
	DefaultAmqpUrl  = "amqp://guest:guest@127.0.0.1:5672"
	DefaultExchange = Exchange{
		Name: "kamyee",
		Type: amqp.ExchangeTopic,
	}
)

type Config struct {
	// broker config
	Username    string
	Password    string
	Host        string
	Port        uint
	VirtualHost string

	// channel config
	Exchange       *Exchange
	PrefetchCount  int
	PrefetchGlobal bool
}

func (config *Config) GetAmqpUrl() string {
	if config.Host == "" {
		return DefaultAmqpUrl
	}

	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.VirtualHost)
}
