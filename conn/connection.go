package conn

type Message interface {
	GetHeaderValue(headerName string) (string, error)
	GetBody() []byte
}

type Connection interface {
	Listen()
	SetNewMsgCallback(f func(Message))
	AckMsg(m Message) error
	NackMsg(m Message) error
	SendResponse(original Message, response Message) error
}
