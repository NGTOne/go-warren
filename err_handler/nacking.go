package err_handler

import(
    "github.com/NGTOne/warren/conn"
)

type NackingErrorHandler struct {
	conn conn.Connection
}

func NewNackingHandler(conn conn.Connection) NackingErrorHandler {
	return NackingErrorHandler{
		conn: conn,
	}
}

func (handler *NackingErrorHandler) ProcessError (
	msg conn.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.NackMessage(msg)
	return e
}
