package err_handler

import (
	"github.com/NGTOne/warren/conn"
)

type nackingHandler struct {
	conn conn.Connection
}

func NewNackingHandler(conn conn.Connection) nackingHandler {
	return nackingHandler{
		conn: conn,
	}
}

func (handler nackingHandler) ProcessErr(
	msg conn.Message,
	e error,
) error {
	// We've already encountered one error; not much we can do if we hit
	// another
	handler.conn.NackMsg(msg)
	return e
}
