package amqp_0_9_1

import (
	"errors"
	"github.com/streadway/amqp"

	"github.com/NGTOne/warren/conn"
)

// streadway/amqp doesn't provide interfaces around a couple of its core types,
// making it impossible to mock them using library builtins
// So we'll create our own instead
type amqpConn interface {
	Channel() (amqpChan, error)
	Close()
}

type amqpChan interface {
	Close()
	Qos(prefetchCount, prefetchSize int, global bool) error
	Ack(deliveryTag uint64, multiple bool) error
	Nack(deliveryTag uint64, multiple bool, requeue bool) error
	Consume(
		queue, consumer string,
		autoAck, exclusive, noLocal, noWait bool,
		args amqp.Table,
	) (<-chan amqp.Delivery, error)

	QueueDeclare(
		name string,
		durable, autoDelete, exclusive, noWait bool,
		args amqp.Table,
	) (amqp.Queue, error)
	ExchangeDeclare(
		name, kind string,
		durable, autoDelete, internal, noWait bool,
		args amqp.Table,
	) error
	QueueBind(
		name, key, exchange string,
		noWait bool,
		args amqp.Table,
	) error
}

// Provides a couple of thin convenience wrappers around streadway's AMQP
// library, so we can plug it into warren in a consistent way
type Connection struct {
	amqpConn amqpConn
	amqpChan amqpChan
	queue    string
}

func NewConn(conn amqpConn) (*Connection, error) {
	defer conn.Close()

	qChan, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	defer qChan.Close()

	err = qChan.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	return &Connection{
		amqpConn: conn,
		amqpChan: qChan,
	}, nil
}

func (c *Connection) AckMsg(m conn.Message) error {
	tag, err := m.GetHeaderValue("DeliveryTag")

	if err != nil {
		return err
	}

	return c.amqpChan.Ack(tag.(uint64), false)
}

func (c *Connection) NackMsg(m conn.Message) error {
	tag, err := m.GetHeaderValue("DeliveryTag")

	if err != nil {
		return err
	}

	return c.amqpChan.Nack(tag.(uint64), false, true)
}

func (c *Connection) Listen(f func(conn.Message)) error {
	if c.queue == "" {
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

	if err != nil {
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
