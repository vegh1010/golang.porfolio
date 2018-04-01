package receiverHelper

import "github.com/streadway/amqp"

//create channel, notify and delivery process
func (self *Receiver) Process(conn *amqp.Connection) {
	defer conn.Close()

	notify := conn.NotifyClose(make(chan *amqp.Error)) //error channel
	forever := make(chan bool)

	msgs := self.Channel(conn)

	go self.Notify(notify, forever)

	go self.Delivery(msgs, forever)

	self.Logger.DebugLog("[*] Waiting for messages...")
	<-forever
}
