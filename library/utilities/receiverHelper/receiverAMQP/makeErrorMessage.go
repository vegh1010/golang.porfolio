package receiver_amqp

import (
	"fmt"
	"errors"
)

func makeErrorMessage(err error, msg string) error {
	return errors.New(fmt.Sprintf("%s: %s", msg, err))
}
