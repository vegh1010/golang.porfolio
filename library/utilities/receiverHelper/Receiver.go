package receiverHelper

import (
	"github.com/streadway/amqp"
	"github.com/vegh1010/golang.porfolio/library/utilities/logHelper"
)

type Receiver struct {
	CredsURL        string
	Logger          logHelper.LogManager
	ProcessDelivery func(d amqp.Delivery)

	//variables for channel's Qos function
	QOSPrefetchCount int
	QOSPrefetchSize  int
	QOSGlobal        bool

	//variables for channel's QueueDeclare function
	DeclareQueueName  string //DR_CDR
	DeclareDurable    bool
	DeclareAutoDelete bool
	DeclareExclusive  bool
	DeclareNoWait     bool
	DeclareArgs       amqp.Table

	//variables for channel's Consume function
	ConsumeConsumer  string
	ConsumeAutoAck   bool
	ConsumeExclusive bool
	ConsumeNoLocal   bool
	ConsumeNoWait    bool
	ConsumeArgs      amqp.Table
}

func (self *Receiver) EnableDebug() {
	self.Logger.EnableDebug()
}

func (self *Receiver) DisableDebug() {
	self.Logger.DisableDebug()
}
