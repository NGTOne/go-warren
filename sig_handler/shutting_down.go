package sig_handler

import(
	"os"
)

type Shutdownable interface{
	ShutDown()
}

type shuttingDownHandler struct{
	// Use a map for easier (and faster) indexing
	signals map[os.Signal]bool
	target  Shutdownable
}

func NewShuttingDownHandler(
	signals []os.Signal,
	target Shutdownable,
) shuttingDownHandler{
	sigMap := make(map[os.Signal]bool)
	for _, sig := range signals {
		sigMap[sig] = true
	}

	return shuttingDownHandler{
		signals: sigMap,
		target: target,
	}
}

func (h shuttingDownHandler) HandleSignals(signals []os.Signal) {
	for _, caughtSig := range signals {
		if _, ok := h.signals[caughtSig]; ok {
			h.target.ShutDown()
			return
		}
	}
}
