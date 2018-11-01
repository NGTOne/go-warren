package err_handler

import (
	"github.com/NGTOne/warren/conn"
)

type ackingHandler struct {
	conn conn.Connection
}

func NewAckingHandler(conn conn.Connection) ackingHandler {
	return ackingHandler{
		conn: conn,
	}
}

func (handler ackingHandler) ProcessErr(
	msg conn.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.AckMsg(msg)
	return e
}
