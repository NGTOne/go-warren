package service_test

import(
	"github.com/NGTOne/warren/service"
	"github.com/NGTOne/warren/test_mocks"

	"testing"
	"github.com/golang/mock/gomock"
	"errors"
)

func TestMissingActionHeader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return(
		"",
		errors.New("Something went wrong!"),
	)

	mockConn.EXPECT().AckMessage(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.Listen()
}

func TestMissingAction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)

	mockConn.EXPECT().AckMessage(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.Listen()
}
