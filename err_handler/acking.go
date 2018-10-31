package err_handler

import(
    "github.com/NGTOne/warren/conn"
)

type AckingErrorHandler struct {
	conn conn.Connection
}

func NewAckingHandler(conn conn.Connection) AckingErrorHandler {
	return AckingErrorHandler{
		conn: conn,
	}
}

func (handler *AckingErrorHandler) ProcessError (
	msg conn.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.AckMessage(msg)
	return e
}
