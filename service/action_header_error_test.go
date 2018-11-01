package service_test

import(
	"github.com/NGTOne/warren/service"
	"github.com/NGTOne/warren/test_mocks"

	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.Listen()
}

func TestNonStringActionHeader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return(-100, nil)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.Listen()
}

func TestMissingAction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	con := service.NewConsumer(mockConn)
	con.Listen()
}

func TestAttemptingToAddSameAsyncActionTwice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	first := test_mocks.NewMockAsynchronousAction(mockCtrl)
	second := test_mocks.NewMockAsynchronousAction(mockCtrl)

	con := service.NewConsumer(mockConn)
	err := con.AddAsyncAction(first, "foo")
	assert.Equal(t, nil, err)

	err = con.AddAsyncAction(second, "foo")
	assert.Equal(
		t,
		errors.New("Action foo already exists in async action list"),
		err,
	)
}

func TestAttemptingToAddSameSyncActionTwice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	first := test_mocks.NewMockSynchronousAction(mockCtrl)
	second := test_mocks.NewMockSynchronousAction(mockCtrl)

	con := service.NewConsumer(mockConn)
	err := con.AddSyncAction(first, "foo")
	assert.Equal(t, nil, err)

	err = con.AddSyncAction(second, "foo")
	assert.Equal(
		t,
		errors.New("Action foo already exists in sync action list"),
		err,
	)
}
