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

func (handler CompositionalHandler) ProcessErr (
	msg conn.Message,
	e error,
) error {
	for _, handler := range handler.handlers{
		e = handler.ProcessErr(msg, e)
	}

	return e
}
