package conn

import (
	"errors"
	"strings"
)

type genericMsg struct {
	headers map[string]interface{}
	body    []byte
}

func NewGenericMsg(headers map[string]interface{}, body []byte) genericMsg {
	return genericMsg{
		body:    body,
		headers: headers,
	}
}

func GenericMsgFromOther(other Message) genericMsg {
	return genericMsg{
		body:    other.GetBody(),
		headers: other.GetAllHeaders(),
	}
}

func (msg genericMsg) GetAllHeaders() map[string]interface{} {
	return msg.headers
}

func (msg genericMsg) GetHeaderValue(headerName string) (interface{}, error) {
	if header, ok := msg.headers[headerName]; ok {
		return header, nil
	}
	return "", errors.New(strings.Join([]string{
		"Header \"",
		headerName,
		"\" not found",
	}, ""))
}

func (msg genericMsg) GetBody() []byte {
	return msg.body
}
