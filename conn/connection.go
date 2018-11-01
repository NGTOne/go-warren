package conn

type Message interface {
	GetAllHeaders() map[string]string
	GetHeaderValue(headerName string) (string, error)
	GetBody() []byte
}

type Connection interface {
	Listen(f func (msg Message))
	AckMsg(m Message) error
	NackMsg(m Message) error
	SendResponse(original Message, response Message) error
}
