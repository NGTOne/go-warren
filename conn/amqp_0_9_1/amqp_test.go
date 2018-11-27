package amqp_0_9_1_test

import(
	"github.com/NGTOne/warren/conn/amqp_0_9_1"
	warren_conn "github.com/NGTOne/warren/conn"

	"github.com/streadway/amqp"

	"testing"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	q_test_mocks "github.com/NGTOne/warren/test_mocks/conn/amqp_0_9_1"
	"github.com/NGTOne/warren/test_mocks"
)

func TestNewConnSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	result, err := amqp_0_9_1.NewConn(mockConn)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestChannelError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)

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

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(
		errors.New("Something went wrong!"),
	)

	result, err := amqp_0_9_1.NewConn(mockConn)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestAckSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("DeliveryTag").Return(uint64(123), nil)
	mockChan.EXPECT().Ack(uint64(123), false).Return(nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.AckMsg(mockMsg)

	assert.Nil(t, err)
}

func TestAckMissingHeader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("DeliveryTag").Return(
		-1,
		errors.New("Something went wrong!"),
	)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.AckMsg(mockMsg)

	assert.Equal(t, errors.New("Something went wrong!"), err)
}

func TestNackSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("DeliveryTag").Return(uint64(123), nil)
	mockChan.EXPECT().Nack(uint64(123), false, true).Return(nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.NackMsg(mockMsg)

	assert.Nil(t, err)
}

func TestNackMissingHeader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("DeliveryTag").Return(
		-1,
		errors.New("Something went wrong!"),
	)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.NackMsg(mockMsg)

	assert.Equal(t, errors.New("Something went wrong!"), err)
}

func TestListenMissingQueue(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
        mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)

        mockConn.EXPECT().Channel().Return(mockChan, nil)
        mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.Listen(func(m warren_conn.Message) {})

	assert.Equal(t, errors.New(
		"Need to create a queue before attempting to listen",
	), err)
}

func TestConsumeError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
        mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)

        mockConn.EXPECT().Channel().Return(mockChan, nil)
        mockChan.EXPECT().Qos(1, 0, false).Return(nil)
	mockChan.EXPECT().QueueDeclare(
		"foo_queue",
		true,
		false,
		false,
		false,
		nil,
	).Return(amqp.Queue{}, nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)
	conn.SetTargetQueue("foo_queue")

	mockChan.EXPECT().Consume(
		"foo_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	).Return(nil, errors.New("Something went wrong!"))

	err := conn.Listen(func(m warren_conn.Message) {})

	assert.Equal(t, errors.New("Something went wrong!"), err)
}
