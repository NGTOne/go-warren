package sig_handler

import(
	"os"
)

type SignalHandler interface {
	HandleSignals(sigs []os.Signal)
}
