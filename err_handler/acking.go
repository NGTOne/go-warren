package err_handler

import(
    "github.com/NGTOne/warren/conn"
)

type AckingHandler struct {
	conn conn.Connection
}

func NewAckingHandler(conn conn.Connection) AckingHandler {
	return AckingHandler{
		conn: conn,
	}
}

func (handler AckingHandler) ProcessErr (
	msg conn.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.AckMessage(msg)
	return e
}
