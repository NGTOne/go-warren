package err_handler

import(
    "github.com/NGTOne/warren/conn"
)

type CompositionalHandler struct {
	handlers []ErrHandler
}

func NewCompositionalHandler(handlers []ErrHandler) CompositionalHandler {
	return CompositionalHandler{
		handlers: handlers,
	}
}

func (handler CompositionalHandler) ProcessError (
	msg conn.Message,
	e error,
) error {
	for _, handler := range handler.handlers{
		e = handler.ProcessError(msg, e)
	}

	return e
}
