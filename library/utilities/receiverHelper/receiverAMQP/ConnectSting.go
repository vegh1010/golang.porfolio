package receiver_amqp

import "fmt"

/*
 * Get amqp string format
 */
func (creds *AMQPCreds) ConnectSting() string {
	return fmt.Sprintf("amqp://%v:%v@%v:%v/", creds.Username, creds.Password, creds.Host, creds.Port)
}
