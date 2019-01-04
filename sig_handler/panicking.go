package sig_handler

import(
	"os"
)

type panickingHandler struct{
	panicMsgs map[os.Signal]string
}

func NewPanickingHandler(panicMsgs map[os.Signal]string) panickingHandler {
	return panickingHandler{
		panicMsgs: panicMsgs,
	}
}

func (h panickingHandler) HandleSignals(sigs []os.Signal) {
	for _, sig := range sigs {
		if msg, ok := h.panicMsgs[sig]; ok {
			panic(msg)
		}
	}
}
