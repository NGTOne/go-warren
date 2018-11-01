package err_handler

import (
	"github.com/NGTOne/warren/conn"
)

type PanickingHandler struct {
	Message string
}

func (handler PanickingHandler) ProcessErr(
	msg conn.Message,
	e error,
) error {
	panic(e)
}
