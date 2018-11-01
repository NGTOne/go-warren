package conn

type Message interface {
	GetAllHeaders() map[string]interface{}
	GetHeaderValue(headerName string) (interface{}, error)
	GetBody() []byte
}

type Connection interface {
	Listen(f func(msg Message)) error
	AckMsg(m Message) error
	NackMsg(m Message) error
	SendResponse(original Message, response Message) error
}
