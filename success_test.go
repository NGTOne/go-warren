package warren_test

import (
	"github.com/NGTOne/warren"
	"github.com/NGTOne/warren/test_mocks"

	"github.com/golang/mock/gomock"
	"testing"
)

func TestSuccessfulAsync(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockAsynchronousAction(mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(nil)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := warren.NewConsumer(mockConn)
	con.AddAsyncAction(mockAction, "foo")

	con.Listen()
}

func TestSuccessfulSync(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockReply := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)
	mockAction := test_mocks.NewMockSynchronousAction(mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)
	mockAction.EXPECT().Process(mockMsg).Return(mockReply, nil)

	mockConn.EXPECT().SendResponse(mockMsg, mockReply).Return(nil)
	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := warren.NewConsumer(mockConn)
	con.AddSyncAction(mockAction, "foo")

	con.Listen()
}
