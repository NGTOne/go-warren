package amqp_0_9_1

import (
	"errors"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetBody(t *testing.T) {
	m := message{
		inner: amqp.Delivery{
			Body: []byte("It's the body!"),
		},
	}

	assert.Equal(t, []byte("It's the body!"), m.GetBody())
}

func TestGetHeaderValue(t *testing.T) {
	m := message{
		inner: amqp.Delivery{
			ContentType:     "text/json",
			ContentEncoding: "base64",
			DeliveryMode:    1,
			Priority:        1,
			CorrelationId:   "f00b4r",
			ReplyTo:         "somewhere",
			Expiration:      "It's gonna expire",
			MessageId:       "Something",
			Timestamp:       time.Unix(100, 0),
			Type:            "Some type",
			UserId:          "Someone",
			AppId:           "An app",
			Headers: map[string]interface{}{
				"some_header": 123,
			},
		},
	}

	tests := []struct {
		header string
		result interface{}
	}{
		{"ContentType", "text/json"},
		{"ContentEncoding", "base64"},
		{"DeliveryMode", uint8(1)},
		{"Priority", uint8(1)},
		{"CorrelationId", "f00b4r"},
		{"ReplyTo", "somewhere"},
		{"Expiration", "It's gonna expire"},
		{"MessageId", "Something"},
		{"Timestamp", time.Unix(100, 0)},
		{"Type", "Some type"},
		{"UserId", "Someone"},
		{"AppId", "An app"},
		{"some_header", 123},
	}

	for _, tt := range tests {
		t.Run(tt.header, func(t *testing.T) {
			result, _ := m.GetHeaderValue(tt.header)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestGetHeaderValueMissingHeader(t *testing.T) {
	m := message{
		inner: amqp.Delivery{
			Headers: map[string]interface{}{
				"some_header": 123,
			},
		},
	}

	_, err := m.GetHeaderValue("some_other_header")

	assert.Equal(
		t,
		errors.New("Header \"some_other_header\" not found"),
		err,
	)
}

func TestGetAllHeaders(t *testing.T) {
	m := message{
		inner: amqp.Delivery{
			ContentType:     "text/json",
			ContentEncoding: "base64",
			DeliveryMode:    1,
			Priority:        1,
			CorrelationId:   "f00b4r",
			ReplyTo:         "somewhere",
			Expiration:      "It's gonna expire",
			MessageId:       "Something",
			Timestamp:       time.Unix(100, 0),
			Type:            "Some type",
			UserId:          "Someone",
			AppId:           "An app",
			Headers: map[string]interface{}{
				"some_header": 123,
			},
		},
	}

	assert.Equal(t, map[string]interface{}{
		"ContentType":     "text/json",
		"ContentEncoding": "base64",
		"DeliveryMode":    uint8(1),
		"Priority":        uint8(1),
		"CorrelationId":   "f00b4r",
		"ReplyTo":         "somewhere",
		"Expiration":      "It's gonna expire",
		"MessageId":       "Something",
		"Timestamp":       time.Unix(100, 0),
		"Type":            "Some type",
		"UserId":          "Someone",
		"AppId":           "An app",
		"some_header":     123,
	}, m.GetAllHeaders())
}
