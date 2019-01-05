package warren

import (
	"os"
	"os/signal"
)

type signalHandler interface {
	HandleSignals(sigs []os.Signal)
}

type signalProcessor struct {
	handler         signalHandler
	caughtSignals   []os.Signal
	catcher         chan os.Signal
	handlingSignals bool
	shutdown        chan bool
}

func newSignalProcessor(handler signalHandler) *signalProcessor {
	p := &signalProcessor{
		handler:         handler,
		caughtSignals:   []os.Signal{},
		catcher:         make(chan os.Signal),
		handlingSignals: true,
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

	// We'll catch as many signals as we can, and let the handler decide
	// which ones it wants to deal with and which ones it wants to ignore
	signal.Stop(p.catcher)
	signal.Notify(
		p.catcher,
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGBUS,
		syscall.SIGCHLD,
		syscall.SIGCLD,
		syscall.SIGCONT,
		syscall.SIGFPE,
		syscall.SIGHUP,
		syscall.SIGILL,
		syscall.SIGINT,
		syscall.SIGIO,
		syscall.SIGIOT,
		syscall.SIGKILL,
		syscall.SIGPIPE,
		syscall.SIGPOLL,
		syscall.SIGPROF,
		syscall.SIGPWR,
		syscall.SIGQUIT,
		syscall.SIGSEGV,
		syscall.SIGSTKFLT,
		syscall.SIGSTOP,
		syscall.SIGSYS,
		syscall.SIGTERM,
		syscall.SIGTRAP,
		syscall.SIGTSTP,
		syscall.SIGTTIN,
		syscall.SIGTTOU,
		syscall.SIGUNUSED,
		syscall.SIGURG,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGVTALRM,
		syscall.SIGWINCH,
		syscall.SIGXCPU,
		syscall.SIGXFSZ,
	)

	return p
}

func (p *signalProcessor) holdSignals() {
	p.handlingSignals = false
}

func (p *signalProcessor) processSignals() {
	p.handler.HandleSignals(p.caughtSignals)
	p.caughtSignals = []os.Signal{}
}

func (p *signalProcessor) stopHoldingSignals() {
	p.handlingSignals = true
}

func (p *signalProcessor) shutDown() {
	p.shutdown <- true
}
