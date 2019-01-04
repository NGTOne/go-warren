package warren

import (
	"os"
	"syscall"
	"errors"
	"github.com/NGTOne/warren/conn"
	"github.com/NGTOne/warren/err_handler"
	"github.com/NGTOne/warren/sig_handler"
	"strings"
)

type consumer struct {
	conn		conn.Connection
	actionHeader	string
	syncActions	map[string]SynchronousAction
	asyncActions	map[string]AsynchronousAction

	sigProcessor      *signalProcessor
	processErrHandler err_handler.ErrHandler
	replyErrHandler   err_handler.ErrHandler
}

func NewConsumer(conn conn.Connection) *consumer {
	sigProcessor := newSignalProcessor(sig_handler.NewPanickingHandler(
		map[os.Signal]string{
			syscall.SIGINT:  "Caught SIGINT",
			syscall.SIGTERM: "Caught SIGTERM",
		},
	))

	return &consumer{
		conn:              conn,
		actionHeader:      "action",
		processErrHandler: err_handler.NewAckingHandler(conn),
		replyErrHandler:   err_handler.NewAckingHandler(conn),
		syncActions:       make(map[string]SynchronousAction),
		asyncActions:      make(map[string]AsynchronousAction),
		sigProcessor:      sigProcessor,
	}
}

func (con *consumer) SetActionHeader(header string) {
	con.actionHeader = header
}

func (con *consumer) SetProcessErrHandler(
	handler err_handler.ErrHandler,
) {
	con.processErrHandler = handler
}

func (con *consumer) SetReplyErrHandler(
	handler err_handler.ErrHandler,
) {
	con.replyErrHandler = handler
}

func (con *consumer) actionAlreadyExists(name string) error {
	if _, alreadyPresent := con.asyncActions[name]; alreadyPresent {
		return errors.New(strings.Join([]string{
			"Action",
			name,
			"already exists in async action list",
		}, " "))
	}

	if _, alreadyPresent := con.syncActions[name]; alreadyPresent {
		return errors.New(strings.Join([]string{
			"Action",
			name,
			"already exists in sync action list",
		}, " "))
	}

	return nil
}

func (con *consumer) AddAsyncAction(
	action AsynchronousAction,
	name string,
) error {
	err := con.actionAlreadyExists(name)
	if err != nil {
		return err
	}

	con.asyncActions[name] = action
	return nil
}

func (con *consumer) AddSyncAction(
	action SynchronousAction,
	name string,
) error {
	err := con.actionAlreadyExists(name)
	if err != nil {
		return err
	}

	con.syncActions[name] = action
	return nil
}

func (con *consumer) Listen() error {
	return con.conn.Listen(func(msg conn.Message) {
		con.processMsg(msg)
	})
}

func (con *consumer) processMsg(msg conn.Message) {
	con.sigProcessor.holdSignals()

	action, err := msg.GetHeaderValue(con.actionHeader)

	if err != nil {
		con.processErrHandler.ProcessErr(msg, err)
		// If we don't know what action to take, we're done here
		return
	}

	if _, ok := action.(string); !ok {
		con.processErrHandler.ProcessErr(
			msg,
			errors.New("Action header was not a string"),
		)
		// We don't know what action to take here, either
		return
	}

	var result conn.Message
	if async, ok := con.asyncActions[action.(string)]; ok {
		err = async.Process(msg)
	} else if sync, ok := con.syncActions[action.(string)]; ok {
		result, err = sync.Process(msg)
	} else {
		con.processErrHandler.ProcessErr(
			msg,
			errors.New(strings.Join([]string{
				"Unknown action",
				action.(string),
			}, " ")),
		)
		return
	}

	if err != nil {
		err = con.processErrHandler.ProcessErr(msg, err)
		return
	}

	if result != nil {
		err = con.conn.SendResponse(msg, result)

		if err != nil {
			con.replyErrHandler.ProcessErr(msg, err)
			return
		}
	}

	err = con.conn.AckMsg(msg)
	if err != nil {
		con.replyErrHandler.ProcessErr(msg, err)
		return
	}
	con.sigProcessor.stopHoldingSignals()
}
