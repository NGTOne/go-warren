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
	// A few special cases for the message properties themselves
	switch(headerName) {
	case "ContentType":
		return m.inner.ContentType, nil
	case "ContentEncoding":
		return m.inner.ContentEncoding, nil
	case "DeliveryMode":
		return m.inner.DeliveryMode, nil
	case "Priority":
		return m.inner.Priority, nil
	case "CorrelationId":
		return m.inner.CorrelationId, nil
	case "ReplyTo":
		return m.inner.ReplyTo, nil
	case "Expiration":
		return m.inner.Expiration, nil
	case "MessageId":
		return m.inner.MessageId, nil
	case "Timestamp":
		return m.inner.Timestamp, nil
	case "Type":
		return m.inner.Type, nil
	case "UserId":
		return m.inner.UserId, nil
	case "AppId":
		return m.inner.AppId, nil
	}

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
	headers := m.inner.Headers

	headers["ContentType"] = m.inner.ContentType
	headers["ContentEncoding"] = m.inner.ContentEncoding
	headers["DeliveryMode"] = m.inner.DeliveryMode
	headers["Priority"] = m.inner.Priority
	headers["CorrelationId"] = m.inner.CorrelationId
	headers["ReplyTo"] = m.inner.ReplyTo
	headers["Expiration"] = m.inner.Expiration
	headers["MessageId"] = m.inner.MessageId
	headers["Timestamp"] = m.inner.Timestamp
	headers["Type"] = m.inner.Type
	headers["UserId"] = m.inner.UserId
	headers["AppId"] = m.inner.AppId

	return headers
}
