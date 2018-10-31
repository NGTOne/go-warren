package service

type ErrorHandler interface {
	ProcessError(msg Message, e error) error
}
