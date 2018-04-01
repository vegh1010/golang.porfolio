package receiver_amqp

import "github.com/streadway/amqp"

/*
 * Ping to check connection to RabbitMQ
 */
func (creds *AMQPCreds) Ping() error {
	conn, err := amqp.Dial(creds.ConnectSting())
	if err != nil {
		return makeErrorMessage(err, "Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	return nil
}
