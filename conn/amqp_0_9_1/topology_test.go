package amqp_0_9_1_test

import (
	"github.com/NGTOne/warren/conn/amqp_0_9_1"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTargetQueueSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := NewMockAMQPChan(mockCtrl)
	mockConn := NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockChan.EXPECT().QueueDeclare(
		"foobar",
		true,
		false,
		false,
		false,
		nil,
	).Return(amqp.Queue{}, nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.SetTargetQueue("foobar")

	assert.Nil(t, err)
}

func TestTargetQueueFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := NewMockAMQPChan(mockCtrl)
	mockConn := NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockChan.EXPECT().QueueDeclare(
		"foobar",
		true,
		false,
		false,
		false,
		nil,
	).Return(amqp.Queue{}, errors.New("Something went wrong!"))

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.SetTargetQueue("foobar")

	assert.NotNil(t, err)
}

func TestCreateAndBindSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := NewMockAMQPChan(mockCtrl)
	mockConn := NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockChan.EXPECT().QueueDeclare(
		"foobar",
		true,
		false,
		false,
		false,
		nil,
	).Return(amqp.Queue{}, nil)
	mockChan.EXPECT().ExchangeDeclare(
		"barbaz",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	).Return(nil)
	mockChan.EXPECT().QueueBind(
		"foobar",
		"",
		"barbaz",
		false,
		nil,
	).Return(nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)
	conn.SetTargetQueue("foobar")

	err := conn.CreateAndBindExchange("barbaz", amqp_0_9_1.Fanout, "")

	assert.Nil(t, err)
}

func TestCreateAndBindNoQueue(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := NewMockAMQPChan(mockCtrl)
	mockConn := NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.CreateAndBindExchange("barbaz", amqp_0_9_1.Fanout, "")

	assert.Equal(t, errors.New(strings.Join([]string{
		"Need to create a queue before attempting to bind it to ",
		"an exchange",
	}, "")), err)
}

func TestCreateAndBindExchangeDeclareFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := NewMockAMQPChan(mockCtrl)
	mockConn := NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockChan.EXPECT().QueueDeclare(
		"foobar",
		true,
		false,
		false,
		false,
		nil,
	).Return(amqp.Queue{}, nil)
	mockChan.EXPECT().ExchangeDeclare(
		"barbaz",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	).Return(errors.New("Something went wrong!"))

	conn, _ := amqp_0_9_1.NewConn(mockConn)
	conn.SetTargetQueue("foobar")

	err := conn.CreateAndBindExchange("barbaz", amqp_0_9_1.Fanout, "")

	assert.Equal(t, errors.New("Something went wrong!"), err)
}

func TestCreateAndBindQueueBindFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := NewMockAMQPChan(mockCtrl)
	mockConn := NewMockAMQPConn(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockChan.EXPECT().QueueDeclare(
		"foobar",
		true,
		false,
		false,
		false,
		nil,
	).Return(amqp.Queue{}, nil)
	mockChan.EXPECT().ExchangeDeclare(
		"barbaz",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	).Return(nil)
	mockChan.EXPECT().QueueBind(
		"foobar",
		"",
		"barbaz",
		false,
		nil,
	).Return(errors.New("Something went wrong!"))

	conn, _ := amqp_0_9_1.NewConn(mockConn)
	conn.SetTargetQueue("foobar")

	err := conn.CreateAndBindExchange("barbaz", amqp_0_9_1.Fanout, "")

	assert.Equal(t, errors.New("Something went wrong!"), err)
}
