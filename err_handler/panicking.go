package err_handler

import(
	"github.com/NGTOne/warren/conn"
)

type PanickingErrorHandler struct {
	Message string
}

func (handler PanickingErrorHandler) ProcessError (
	msg conn.Message,
	e error,
) error {
	panic(e)
}
