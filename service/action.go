package service

type SynchronousAction interface {
	Process(msg Message) (Message, error)
}

type AsynchronousAction interface {
	Process(msg Message) error
}
