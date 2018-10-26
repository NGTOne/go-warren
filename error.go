package warren

type ErrorHandler interface {
	ProcessError(msg Message, e error)
}
