package warren

import (
	"github.com/NGTOne/warren/sig_handler"
	"os"
	"os/signal"
)

type signalProcessor struct {
	handler         sig_handler.SignalHandler
	caughtSignals   []os.Signal
	catcher         chan os.Signal
	handlingSignals bool
	shutdown        chan bool
}

func newSignalProcessor() *signalProcessor {
	p := &signalProcessor{
		handler:         nil,
		caughtSignals:   []os.Signal{},
		catcher:         make(chan os.Signal),
		handlingSignals: false,
		shutdown:        make(chan bool),
	}

	go func() {
		for {
			select {
			case sig := <-p.catcher:
				p.caughtSignals = append(p.caughtSignals, sig)
			case <-p.shutdown:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-p.shutdown:
				return
			default:
				if p.handlingSignals {
					p.processSignals()
				}
			}
		}
	}()

	return p
}

func (p *signalProcessor) setHandler(handler sig_handler.SignalHandler) {
	p.handler = handler
}

func (p *signalProcessor) setTargetSignals(signals []os.Signal) {
	signal.Stop(p.catcher)
	signal.Notify(p.catcher, signals...)
}

func (p *signalProcessor) holdSignals() {
	p.handlingSignals = false
}

func (p *signalProcessor) processSignals() {
	p.handler.HandleSignals(p.caughtSignals)
}

func (p *signalProcessor) stopHoldingSignals() {
	p.handlingSignals = true
}

func (p *signalProcessor) shutDown() {
	p.shutdown <- true
}
