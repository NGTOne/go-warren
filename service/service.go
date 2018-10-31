package service

import (
	"errors"
	"strings"
	"github.com/NGTOne/warren/err_handler"
	"github.com/NGTOne/warren/conn"
)

type consumer struct {
	conn conn.Connection
	actionHeader string
	syncActions map[string]SynchronousAction
	asyncActions map[string]AsynchronousAction

	processErrHandler ErrorHandler
	replyErrHandler ErrorHandler
}

func NewConsumer(conn conn.Connection) *consumer {
	return &consumer{
		conn: conn,
		actionHeader: "action",
		processErrHandler: err_handler.NewAckingHandler(conn),
		replyErrHandler: err_handler.NewAckingHandler(conn),
		syncActions: make(map[string]SynchronousAction),
		asyncActions: make(map[string]AsynchronousAction),
	}
}

func (con *consumer) SetActionHeader (header string) {
	con.actionHeader = header
}

func (con *consumer) SetProcessErrorHandler(handler ErrorHandler) {
	con.processErrHandler = handler
}

func (con *consumer) SetReplyErrorHandler(handler ErrorHandler) {
	con.replyErrHandler = handler
}

func (con *consumer) actionAlreadyExists (name string) error {
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
	if (err != nil) {
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
	if (err != nil) {
		return err
	}

	con.syncActions[name] = action
	return nil
}

func (con *consumer) Listen() {
	con.conn.SetNewMessageCallback(func (msg conn.Message) {
		con.processMsg(msg)
	})

	con.conn.Listen()
}

func (con *consumer) processMsg(msg conn.Message) {
	action, err := msg.GetHeaderValue(con.actionHeader)

	if (err != nil) {
		err = con.processErrHandler.ProcessError(msg, err)
		// If we don't know what action to take, we're done here
		if (err != nil) {
			return
		}
	}

	var result conn.Message
	if async, ok := con.asyncActions[action]; ok {
		err = async.Process(msg)
	} else if sync, ok := con.syncActions[action]; ok {
		result, err = sync.Process(msg)
	} else {
		err = con.processErrHandler.ProcessError(
			msg,
			errors.New(strings.Join([]string{
				"Unknown action",
				action,
			}, " ")),
		)
		if (err != nil) {
			return
		}
	}

	if (err != nil) {
		err = con.processErrHandler.ProcessError(msg, err)
		if (err != nil) {
			return
		}
	}

	if (result != nil) {
		err = con.conn.SendResponse(msg, result)

		if (err != nil) {
			err = con.replyErrHandler.ProcessError(msg, err)
			if (err != nil) {
				return
			}
		}
	}

	err = con.conn.AckMessage(msg)
	if (err != nil) {
		err = con.replyErrHandler.ProcessError(msg, err)
		if (err != nil) {
			return
		}
	}
}
