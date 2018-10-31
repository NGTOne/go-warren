package service_test

import(
	"github.com/NGTOne/warren/service"
	"github.com/NGTOne/warren/test_mocks"

	"testing"
	"github.com/golang/mock/gomock"
	"errors"
)

func TestUnsuccessfulProcessing(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockAsynchronousAction(mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(
		errors.New("Something went wrong!"),
	)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.AddAsyncAction(mockAction, "foo")

	con.Listen()
}

func TestUnsuccessfulReply(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockReply := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockSynchronousAction(mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(mockReply, nil)

	mockConn.EXPECT().SendResponse(mockMsg, mockReply).Return(
		errors.New("Something went wrong!"),
	)
	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.AddSyncAction(mockAction, "foo")

	con.Listen()
}

func TestUnsuccessfulAck(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockAsynchronousAction(mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(nil)

	mockConn.EXPECT().AckMsg(mockMsg).Return(
		errors.New("Something went wrong!"),
	)
	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.AddAsyncAction(mockAction, "foo")

	con.Listen()
}
