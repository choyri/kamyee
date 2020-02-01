package rabbitmq

import (
	"github.com/streadway/amqp"
	"regexp"
)

type Options struct {
	URL           string
	Exchange      exchange
	PrefetchCount uint16
}

type Option func(*Options)

func URL(v string) Option {
	return func(o *Options) {
		if regexp.MustCompile("^amqp(s)?://.*").MatchString(v) {
			o.URL = v
		} else {
			o.URL = DefaultAMQPURL
		}
	}
}

func ExchangeName(v string) Option {
	return func(o *Options) {
		if len(o.Exchange.Name) == 0 {
			o.Exchange = DefaultExchange
		}
		o.Exchange.Name = v
	}
}

func ExchangeType(v string) Option {
	return func(o *Options) {
		if len(o.Exchange.Kind) == 0 {
			o.Exchange = DefaultExchange
		}
		o.Exchange.Kind = v
	}
}

func PrefetchCount(v uint16) Option {
	return func(o *Options) {
		o.PrefetchCount = v
	}
}

func DelayedExchange(name string, xDelayedType ...string) Option {
	kind := amqp.ExchangeTopic

	if len(xDelayedType) > 0 {
		kind = xDelayedType[0]
	}

	return func(o *Options) {
		o.Exchange = exchange{
			Name: name,
			Kind: delayedExchangeType,
			Args: amqp.Table{
				delayedExchangeArgsKey: kind,
			},
		}
	}
}
