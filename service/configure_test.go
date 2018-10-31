package service_test

import(
	"github.com/NGTOne/warren/service"
	"github.com/NGTOne/warren/test_mocks"

	"testing"
	"github.com/golang/mock/gomock"
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
