package service_test

import(
	"github.com/NGTOne/warren/service"
	"github.com/NGTOne/warren/test_mocks"

	"testing"
	"github.com/golang/mock/gomock"
	"errors"
)

func TestDifferentActionHeader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockAsynchronousAction(mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("foobar").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(nil)

	mockConn.EXPECT().AckMessage(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.AddAsyncAction(mockAction, "foo")
	con.SetActionHeader("foobar")

	con.Listen()
}

func TestDifferentProcessHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockAsynchronousAction(mockCtrl)

	err := errors.New("Something went wrong!")

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(err)

	mockHandler := test_mocks.NewMockErrHandler(mockCtrl)
	mockHandler.EXPECT().ProcessError(mockMsg, err).Return(err)

	con := service.NewConsumer(mockConn)
	con.AddAsyncAction(mockAction, "foo")
	con.SetProcessErrHandler(mockHandler)

	con.Listen()
}

func TestDifferentReplyHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockSynchronousAction(mockCtrl)

	err := errors.New("Something went wrong!")

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(mockMsg, nil)

	mockHandler := test_mocks.NewMockErrHandler(mockCtrl)
	mockHandler.EXPECT().ProcessError(mockMsg, err).Return(err)

	mockConn.EXPECT().SendResponse(mockMsg, mockMsg).Return(err)

	con := service.NewConsumer(mockConn)
	con.AddSyncAction(mockAction, "foo")
	con.SetReplyErrHandler(mockHandler)

	con.Listen()
}
