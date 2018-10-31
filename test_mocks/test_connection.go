package test_mocks

import(
	"github.com/NGTOne/warren/conn"

	"errors"
)

type TestConnection struct{
	callback func(conn.Message)

	Message conn.Message
	AckError bool
	NackError bool
	ReplyError bool
}

func (con *TestConnection) Listen() {
	con.callback(con.Message)
}

func (con *TestConnection) SetNewMessageCallback(f func(conn.Message)) {
	con.callback = f
}

func (con *TestConnection) AckMessage(m conn.Message) error {
	if (con.AckError) {
		return errors.New("Something went wrong!")
	}
	return nil
}

func (con *TestConnection) NackMessage(m conn.Message) error {
	if (con.NackError) {
		return errors.New("Something went wrong!")
	}
	return nil
}

func (con *TestConnection) SendResponse(
	original conn.Message,
	response conn.Message,
) error {
	if (con.ReplyError) {
		return errors.New("Something went wrong!")
	}
	return nil
}
