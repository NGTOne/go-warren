package amqp_0_9_1

import(
	"github.com/streadway/amqp"
	"errors"

	"github.com/NGTOne/warren/conn"
)

// Provides a couple of thin convenience wrappers around streadway's AMQP
// library, so we can plug it into warren in a consistent way
type Connection struct {
	amqpConn *amqp.Connection
	amqpChan *amqp.Channel
	queue string
}

func NewConn(url string) (*Connection, error) {
	amqpConn, err := amqp.Dial(url)
	if (err != nil) {
		return nil, err
	}

	defer amqpConn.Close()

	amqpChan, err := amqpConn.Channel()
	if (err != nil) {
		return nil, err
	}

	defer amqpChan.Close()

	err = amqpChan.Qos(1, 0, false)
	if (err != nil) {
		return nil, err
	}

	return &Connection{
		amqpConn: amqpConn,
		amqpChan: amqpChan,
	}, nil
}

func (c *Connection) AckMsg(m conn.Message) error {
	tag, err := m.GetHeaderValue("DeliveryTag")

	if (err != nil) {
		return err
	}

	return c.amqpChan.Ack(tag.(uint64), false)
}

func (c *Connection) NackMsg(m conn.Message) error {
	tag, err := m.GetHeaderValue("DeliveryTag")

	if (err != nil) {
		return err
	}

	return c.amqpChan.Nack(tag.(uint64), false, true)
}

func (c *Connection) Listen(f func(conn.Message)) error {
	if (c.queue == "") {
		return errors.New(
                        "Need to create a queue before attempting to listen",
                )
	}

	msgs, err := c.amqpChan.Consume(
		c.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if (err != nil) {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			f(message{inner: d})
		}
	}()

	<-forever

	return nil
}
