package warren

import(
	"os"
	"os/signal"
	"github.com/NGTOne/warren/sig_handler"
)

type signalProcessor struct {
	handler		sig_handler.SignalHandler
	targetSignals	[]os.Signal
	caughtSignals	[]os.Signal
	catcher		chan os.Signal
}

func newProcessor() *signalProcessor {
	p := &signalProcessor{
		handler:	nil,
		targetSignals:	[]os.Signal{},
		caughtSignals:	[]os.Signal{},
		catcher:	make(chan os.Signal),
	}

	go func() {
		for {
			sig := <-p.catcher
			p.caughtSignals = append(p.caughtSignals, sig)
		}
	}()

	return p
}

func (p *signalProcessor) setHandler (handler sig_handler.SignalHandler) {
	p.handler = handler
}

func (p *signalProcessor) setTargetSignals(signals []os.Signal) {
	p.targetSignals = signals
}

func (p *signalProcessor) holdSignals() {
	signal.Notify(p.catcher, p.targetSignals...)
}

func (p *signalProcessor) processSignals() {
	p.handler.HandleSignals(p.caughtSignals)
}

func (p *signalProcessor) stopHoldingSignals() {
	signal.Stop(p.catcher)
}
