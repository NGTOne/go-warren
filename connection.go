package warren

type Message interface {
	GetHeaderValue(headerName string) (string, error)
	GetBody() []byte
}

type Connection interface {
	Listen()
	SetNewMessageCallback(f func(Message))
	AckMessage(m Message) error
	NackMessage(m Message) error
	SendResponse(original Message, response Message) error
}
