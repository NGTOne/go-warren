package err_handler

import(
	"github.com/NGTOne/warren/conn"
)

type PanickingHandler struct {
	Message string
}

func (handler PanickingHandler) ProcessError (
	msg conn.Message,
	e error,
) error {
	panic(e)
}
