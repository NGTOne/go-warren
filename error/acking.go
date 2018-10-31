package error

import(
    "github.com/NGTOne/warren/service"
)

type AckingErrorHandler struct {
	conn service.Connection
}

func NewAckingHandler(conn service.Connection) AckingErrorHandler {
	return AckingErrorHandler{
		conn: conn,
	}
}

func (handler *AckingErrorHandler) ProcessError (
	msg service.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.AckMessage(msg)
	return e
}
