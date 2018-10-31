package err_handler

import(
	"github.com/NGTOne/warren/conn"
)

type ErrHandler interface {
	ProcessError(msg conn.Message, e error) error
}
