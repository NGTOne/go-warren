package amqp_0_9_1

import(
	"github.com/streadway/amqp"
)

type message struct{
	inner amqp.Delivery
}

func newMessage(d amqp.Delivery) message {
	return message{
		inner: d,
	}
}
