package amqp_0_9_1_test

import(
	"github.com/NGTOne/warren/conn/amqp_0_9_1"

	"testing"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	test_mocks "github.com/NGTOne/warren/test_mocks/conn/amqp_0_9_1"
)

func TestNewConnSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := test_mocks.NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	result, err := amqp_0_9_1.NewConn(mockConn)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestChannelError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConn := test_mocks.NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(
		nil,
		errors.New("Something went wrong!"),
	)

	result, err := amqp_0_9_1.NewConn(mockConn)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestQosError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := test_mocks.NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(
		errors.New("Something went wrong!"),
	)

	result, err := amqp_0_9_1.NewConn(mockConn)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}
