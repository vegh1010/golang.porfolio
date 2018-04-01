package receiverHelper

import "github.com/streadway/amqp"

//process messages on queue
func (self *Receiver) Delivery(msgs <-chan amqp.Delivery, forever chan bool) {
	for d := range msgs {
		self.Logger.DebugLog("[QUEUE]", "Received Message with Body: ", string(d.Body))

		//run delivery process function
		self.ProcessDelivery(d)
	}
}
