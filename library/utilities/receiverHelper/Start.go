package receiverHelper

import (
	"github.com/streadway/amqp"
	"net"
	"time"
)

//start rabbitmq dial
func (self *Receiver) Start() {
	for {
		self.Logger.DebugLog("Dial AMQP")
		conn, err := amqp.DialConfig(self.CredsURL, amqp.Config{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, 30*time.Second)
			},
		})
		FailOnError(err, "Failed to connect to RabbitMQ")

		self.Process(conn)
	}
}
