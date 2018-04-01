package receiverHelper

import "github.com/streadway/amqp"

//establish channel and queue, return chan Delivery (messages)
func (self *Receiver) Channel(conn *amqp.Connection) (<-chan amqp.Delivery) {
	self.Logger.DebugLog("Get Channel")
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	self.Logger.DebugLog("Set Channel Qos")
	err = ch.Qos(
		self.QOSPrefetchCount, // prefetch count
		self.QOSPrefetchSize,  // prefetch size
		self.QOSGlobal,        // global
	)
	FailOnError(err, "Failed to set QoS")

	self.Logger.DebugLog("Declare Channel's Queue")
	q, err := ch.QueueDeclare(
		self.DeclareQueueName,  // name
		self.DeclareDurable,    // durable
		self.DeclareAutoDelete, // delete when usused
		self.DeclareExclusive,  // exclusive
		self.DeclareNoWait,     // no-wait
		self.DeclareArgs,       // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	self.Logger.DebugLog("Get Channel's Messages")
	msgs, err := ch.Consume(
		q.Name,                // queue
		self.ConsumeConsumer,  // consumer
		self.ConsumeAutoAck,   // auto-ack
		self.ConsumeExclusive, // exclusive
		self.ConsumeNoLocal,   // no-local
		self.ConsumeNoWait,    // no-wait
		self.ConsumeArgs,      // args
	)
	FailOnError(err, "Failed to register a consumer")

	return msgs
}
