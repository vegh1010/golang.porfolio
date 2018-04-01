package receiverHelper

import "github.com/streadway/amqp"

//release forever chan if notify error received
func (self *Receiver) Notify(notify chan *amqp.Error, forever chan bool) {
	errN := <-notify
	self.Logger.DebugLog("Notify Error: " + errN.Error())
	forever <- true
}
