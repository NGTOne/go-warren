package amqp_0_9_1

import(
	"github.com/streadway/amqp"
	"errors"
	"strings"
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

	return &Connection{
		amqpConn: amqpConn,
		amqpChan: amqpChan,
	}, nil
}

func (c *Connection) SetTargetQueue(name string) error {
	// We're gonna make a few assumptions here - since this lib is mainly
	// meant to take the place of "traditional" HTTP-based microservices,
	// we'll make the queue stick around forever
	_, err := c.amqpChan.QueueDeclare(name, true, false, false, false, nil)

	if (err != nil) {
		c.queue = name
	}

	return err
}

type ExchangeType string

const(
	Direct	ExchangeType = "direct"
	Fanout	ExchangeType = "fanout"
	Topic	ExchangeType = "topic"
	Headers	ExchangeType = "headers"
)

func (c *Connection) CreateAndBindExchange(
	name string,
	exchType ExchangeType,
	routingKey string,
) error {
	if (c.queue == "") {
		return errors.New(strings.Join([]string{
			"Need to create a queue before attempting to bind ",
			"it to an exchange",
		}, ""))
	}

	// Like before, we're gonna make a few basic assumptions here just
	// to make life simple
	err := c.amqpChan.ExchangeDeclare(
		name,
		string(exchType),
		true,
		false,
		false,
		false,
		nil,
	)

	if (err != nil) {
		return err
	}

	err = c.amqpChan.QueueBind(c.queue, routingKey, name, false,nil)
	if (err != nil) {
		return err
	}

	return nil
}
