package amqp_0_9_1

import (
	"errors"
	"github.com/streadway/amqp"
	"strings"
)

type message struct {
	inner amqp.Delivery
}

func newMessage(d amqp.Delivery) message {
	return message{
		inner: d,
	}
}

func (m message) GetBody() []byte {
	return m.inner.Body
}

// I'm not _100%_ happy with how this works because the headers can be
// recursive, but I'm not sure how best to deal with it for now
func (m message) GetHeaderValue(headerName string) (interface{}, error) {
	if value, ok := m.inner.Headers[headerName]; ok {
		return value, nil
	}

	return nil, errors.New(strings.Join([]string{
		"Header \"",
		headerName,
		"\" not found",
	}, ""))
}

func (m message) GetAllHeaders() map[string]interface{} {
	return m.inner.Headers
}
