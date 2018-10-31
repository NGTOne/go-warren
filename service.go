package warren

import (
	"errors"
	"strings"
)

type Consumer struct {
	conn Connection
	actionHeader string
	syncActions map[string]SynchronousAction
	asyncActions map[string]AsynchronousAction

	processErrHandler ErrorHandler
	replyErrHandler ErrorHandler
}

func NewConsumer(conn Connection) *Consumer {
	return &Consumer{
		conn: conn,
	}
}

func (con *Consumer) SetActionHeader (header string) {
	con.actionHeader = header
}

func (con *Consumer) actionAlreadyExists (name string) error {
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

func (con *Consumer) AddAsyncAction(
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

func (con *Consumer) AddSyncAction(
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

func (con *Consumer) Listen() {
	con.conn.SetNewMessageCallback(func (msg Message) {
		con.processMsg(msg)
	})

	con.conn.Listen()
}

func (con *Consumer) cleanUpAfterError(msg Message) {
	// Not much point in checking the error here; if we fail to ack
	// the message we're kinda screwed anyways
	con.conn.AcknowledgeMessage(msg)
}

func (con *Consumer) processMsg(msg Message) {
	action, err := msg.GetHeaderValue(con.actionHeader)

	if (err != nil) {
		err = con.processErrHandler.ProcessError(msg, err)
		con.cleanUpAfterError(msg)
		// If we don't know what action to take, we're done here
		if (err != nil) {
			return
		}
	}

	var result Message
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
		con.cleanUpAfterError(msg)
		if (err != nil) {
			return
		}
	}

	if (err != nil) {
		err = con.processErrHandler.ProcessError(
			msg,
			err,
		)
		con.cleanUpAfterError(msg)
		if (err != nil) {
			return
		}
	}

	if (result != nil) {
		err = con.conn.SendResponse(msg, result)
	}
}
