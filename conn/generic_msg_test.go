package conn_test

import (
	"github.com/NGTOne/warren/conn"

	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessage(t *testing.T) {
	body := []byte("This is the body.")
	headers := map[string]interface{}{
		"foo":    "bar",
		"action": "consequence",
	}

	msg := conn.NewGenericMsg(headers, body)

	assert.Equal(t, body, msg.GetBody())
	assert.Equal(t, headers, msg.GetAllHeaders())
}

func TestMissingHeader(t *testing.T) {
	body := []byte{}
	headers := map[string]interface{}{}

	msg := conn.NewGenericMsg(headers, body)

	result, err := msg.GetHeaderValue("action")

	assert.Equal(t, "", result)
	assert.Equal(t, errors.New("Header \"action\" not found"), err)
}

func TestPresentHeader(t *testing.T) {
	body := []byte{}
	headers := map[string]interface{}{
		"action": "consequence",
	}

	msg := conn.NewGenericMsg(headers, body)

	result, err := msg.GetHeaderValue("action")
	assert.Equal(t, "consequence", result)
	assert.Nil(t, err)
}

func TestFromOther(t *testing.T) {
	body := []byte{}
	headers := map[string]interface{}{
		"action": "consequence",
	}

	msg := conn.NewGenericMsg(headers, body)

	result := conn.GenericMsgFromOther(msg)

	assert.Equal(t, body, result.GetBody())
	assert.Equal(t, headers, result.GetAllHeaders())
}
