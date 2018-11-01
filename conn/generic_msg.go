package conn

import(
	"strings"
	"errors"
)

type GenericMsg struct {
	headers map[string]string
	body []byte
}

func NewGenericMsg(headers map[string]string, body []byte) GenericMsg {
	return GenericMsg{
		body: body,
		headers: headers,
	}
}

func GenericMsgFromOther(other Message) GenericMsg {
	return GenericMsg{
		body: other.GetBody(),
		headers: other.GetAllHeaders(),
	}
}

func (msg GenericMsg) GetAllHeaders() map[string]string {
	return msg.headers
}

func (msg GenericMsg) GetHeaderValue(headerName string) (string, error) {
	if header, ok := msg.headers[headerName]; ok {
		return header, nil
	}
	return "", errors.New(strings.Join([]string{
		"Header \"",
		headerName,
		"\" not found",
	}, ""))
}

func (msg GenericMsg) GetBody() []byte {
	return msg.body
}
