package service

import(
	"github.com/NGTOne/warren/conn"
)

type ErrorHandler interface {
	ProcessError(msg conn.Message, e error) error
}
