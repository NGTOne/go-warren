package amqp_0_9_1_test

import(
	"github.com/NGTOne/warren/conn/amqp_0_9_1"

	"testing"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/streadway/amqp"
	q_test_mocks "github.com/NGTOne/warren/test_mocks/conn/amqp_0_9_1"
	"github.com/NGTOne/warren/test_mocks"
)

func TestSendResponseSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockRes := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("ReplyTo").Return("inbox", nil)
	mockMsg.EXPECT().GetHeaderValue("CorrelationId").Return("f00b4r", nil)

	mockRes.EXPECT().GetBody().Return([]byte("This is a response"))
	mockRes.EXPECT().GetAllHeaders().Return(map[string]interface{}{
		"foo": "bar",
		"baz": -15,
	})

	mockChan.EXPECT().Publish("", "inbox", false, false, amqp.Publishing{
		ContentType: "text/plain",
		CorrelationId: "f00b4r",
		Headers: map[string]interface{}{
			"foo": "bar",
			"baz": -15,
		},
		Body: []byte("This is a response"),
	})

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.SendResponse(mockMsg, mockRes)

	assert.Nil(t, err)
}

func TestNoCorrID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockRes := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("CorrelationId").Return(
		"",
		errors.New("Something went wrong!"),
	)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.SendResponse(mockMsg, mockRes)

	assert.Equal(t, errors.New("Something went wrong!"), err)
}

func TestNonStringCorrID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockChan := q_test_mocks.NewMockAMQPChan(mockCtrl)
	mockConn := q_test_mocks.NewMockAMQPConn(mockCtrl)
	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockRes := test_mocks.NewMockMessage(mockCtrl)

	mockConn.EXPECT().Channel().Return(mockChan, nil)
	mockChan.EXPECT().Qos(1, 0, false).Return(nil)

	mockMsg.EXPECT().GetHeaderValue("CorrelationId").Return(42, nil)

	conn, _ := amqp_0_9_1.NewConn(mockConn)

	err := conn.SendResponse(mockMsg, mockRes)

	assert.Equal(t, errors.New("Correlation ID is not a string"), err)
}
