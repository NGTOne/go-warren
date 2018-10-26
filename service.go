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
}

func NewConsumer(conn Connection) *Consumer {
	return &Consumer{
		conn: conn,
	}
}

func (conn *Consumer) SetActionHeader (header string) {
	conn.actionHeader = header
}

func (conn *Consumer) actionAlreadyExists (name string) error {
	if _, alreadyPresent := conn.asyncActions[name]; alreadyPresent {
		return errors.New(strings.Join([]string{
			"Action",
			name,
			"already exists in async action list",
		}, " "))
	}

	if _, alreadyPresent := conn.syncActions[name]; alreadyPresent {
		return errors.New(strings.Join([]string{
			"Action",
			name,
			"already exists in sync action list",
		}, " "))
	}

	return nil
}

func (conn *Consumer) AddAsyncAction(
	action AsynchronousAction,
	name string,
) error {
	err := conn.actionAlreadyExists(name)
	if (err != nil) {
		return err
	}

	conn.asyncActions[name] = action
	return nil
}

func (conn *Consumer) AddSyncAction(
	action SynchronousAction,
	name string,
) error {
	err := conn.actionAlreadyExists(name)
	if (err != nil) {
		return err
	}

	conn.syncActions[name] = action
	return nil
}
