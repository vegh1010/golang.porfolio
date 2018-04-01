package receiver_amqp

import (
	"sync"
)

/*
 * struct: AMQP Credentials properties
 */
type AMQPCreds struct {
	Username string
	Password string
	Host     string
	Port     string
	WG       sync.WaitGroup
}

// NewValidator creates a new Validator
func (creds *AMQPCreds) New() *AMQPCreds {
	return &AMQPCreds{
		Username: creds.Username,
		Password: creds.Password,
		Host:     creds.Host,
		Port:     creds.Port,
	}
}
