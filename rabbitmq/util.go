package kyrabbitmq

import (
	"errors"
	"fmt"
)

func WrapErr(format string, err error) error {
	return errors.New(fmt.Sprintf("[kyRabbitMQ] %s: %s", format, err.Error()))
}
