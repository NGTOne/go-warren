package err_handler

import (
	"github.com/NGTOne/warren/conn"
)

type compositionalHandler struct {
	handlers []ErrHandler
}

func NewCompositionalHandler(handlers []ErrHandler) compositionalHandler {
	return compositionalHandler{
		handlers: handlers,
	}
}

func (handler compositionalHandler) ProcessErr(
	msg conn.Message,
	e error,
) error {
	for _, handler := range handler.handlers {
		e = handler.ProcessErr(msg, e)
	}

	return e
}
