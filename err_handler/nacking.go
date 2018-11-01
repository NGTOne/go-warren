package err_handler

import (
	"github.com/NGTOne/warren/conn"
)

type NackingHandler struct {
	conn conn.Connection
}

func NewNackingHandler(conn conn.Connection) NackingHandler {
	return NackingHandler{
		conn: conn,
	}
}

func (handler NackingHandler) ProcessErr(
	msg conn.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.NackMsg(msg)
	return e
}
