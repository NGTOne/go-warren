package error

import(
    "github.com/NGTOne/warren/service"
)

type NackingErrorHandler struct {
	conn service.Connection
}

func NewNackingHandler(conn service.Connection) NackingErrorHandler {
	return NackingErrorHandler{
		conn: conn,
	}
}

func (handler *NackingErrorHandler) ProcessError (
	msg service.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.NackMessage(msg)
	return e
}
