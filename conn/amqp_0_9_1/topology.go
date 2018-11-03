package amqp_0_9_1

import (
	"errors"
	"strings"
)

func (c *connection) SetTargetQueue(name string) error {
	// We're gonna make a few assumptions here - since this lib is mainly
	// meant to take the place of "traditional" HTTP-based microservices,
	// we'll make the queue stick around forever
	_, err := c.qChan.QueueDeclare(name, true, false, false, false, nil)

	if err != nil {
		c.queue = name
	}

	return err
}

type ExchangeType string

const (
	Direct  ExchangeType = "direct"
	Fanout  ExchangeType = "fanout"
	Topic   ExchangeType = "topic"
	Headers ExchangeType = "headers"
)

func (c *connection) CreateAndBindExchange(
	name string,
	exchType ExchangeType,
	routingKey string,
) error {
	if c.queue == "" {
		return errors.New(strings.Join([]string{
			"Need to create a queue before attempting to bind ",
			"it to an exchange",
		}, ""))
	}

	// Like before, we're gonna make a few basic assumptions here just
	// to make life simple
	err := c.qChan.ExchangeDeclare(
		name,
		string(exchType),
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = c.qChan.QueueBind(c.queue, routingKey, name, false, nil)
	if err != nil {
		return err
	}

	return nil
}
