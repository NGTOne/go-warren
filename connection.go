package warren

type Message interface {
	GetHeaderValue(headerName string) (string, error)
	GetBody() []byte
}

type Connection interface {
	Listen()
	SetNewMessageCallback(f func(Message))
	AcknowledgeMessage(m Message) error
	SendResponse(original Message, response Message) error
}
