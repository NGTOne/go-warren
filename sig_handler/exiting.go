package sig_handler

import(
	"os"
	"runtime"
)

struct ExitingHandler{
	exitCodes: map[os.Signal]int
}

func NewExitingHandler(exitCodes map[os.Signal]int) ExitingHandler {
	return ExitingHandler{
		exitCodes:	exitCodes
	}
}

func (h ExitingHandler) handleSignals(sigs []os.Signal) {
	for _, sig := range sigs {
		if code, ok := h.exitCodes[sig] {
			runtime.Goexit(code)
		}
	}
}
