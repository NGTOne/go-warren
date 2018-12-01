package warren

import (
	"github.com/NGTOne/warren/conn"
)

type SynchronousAction interface {
	Process(msg conn.Message) (conn.Message, error)
}

type AsynchronousAction interface {
	Process(msg conn.Message) error
}
