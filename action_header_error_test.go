package warren_test

import (
	"github.com/NGTOne/warren"
	"github.com/NGTOne/warren/test_mocks"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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
	mockConn.EXPECT().Disconnect()

	con := warren.NewConsumer(mockConn)
	defer con.ShutDown()
	con.Listen()
}

func TestNonStringActionHeader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return(-100, nil)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)
	mockConn.EXPECT().Disconnect()

	con := warren.NewConsumer(mockConn)
	defer con.ShutDown()
	con.Listen()
}

func TestMissingAction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockMsg.EXPECT().GetHeaderValue("action").Return("foo", nil)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)
	mockConn.EXPECT().Disconnect()

	con := warren.NewConsumer(mockConn)
	defer con.ShutDown()
	con.Listen()
}

func TestAttemptingToAddSameAsyncActionTwice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	first := test_mocks.NewMockAsynchronousAction(mockCtrl)
	second := test_mocks.NewMockAsynchronousAction(mockCtrl)

	con := warren.NewConsumer(mockConn)
	defer con.ShutDown()
	mockConn.EXPECT().Disconnect()

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

	con := warren.NewConsumer(mockConn)
	defer con.ShutDown()
	mockConn.EXPECT().Disconnect()

	err := con.AddSyncAction(first, "foo")
	assert.Equal(t, nil, err)

	err = con.AddSyncAction(second, "foo")
	assert.Equal(
		t,
		errors.New("Action foo already exists in sync action list"),
		err,
	)
}
