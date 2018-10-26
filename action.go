package warren

type SynchronousAction interface {
	Process(msg Message) Message
}

type AsynchronousAction interface {
	Process(msg Message)
}
