package warren

type Message interface {
	GetHeaderValue(headerName string) (string, error)
	GetBody() []byte
}

type Connection interface {
	Listen()
	SetNewMessageCallback(f func(Message))
	AcknowledgeMessage(m Message)
	SendResponse(original Message, response Message)
}
