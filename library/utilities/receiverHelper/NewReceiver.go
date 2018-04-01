package receiverHelper

import (
	"github.com/streadway/amqp"
)

//create new instance of QueueReceiver
func NewReceiver(CredsURL string, QueueName string, Process func(d amqp.Delivery)) (*Receiver) {
	manager := Receiver{}
	manager.CredsURL = CredsURL
	manager.DeclareQueueName = QueueName
	manager.ProcessDelivery = Process
	manager.QOSPrefetchCount = 1
	manager.Logger.DevMode()
	manager.Logger.EnableDebug()

	return &manager
}
