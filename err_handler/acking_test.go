package err_handler_test

import (
	"github.com/NGTOne/warren/err_handler"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAckingHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := NewMockMessage(mockCtrl)
	mockConn := NewMockConnection(mockCtrl)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

	err := errors.New("Something went wrong!")

	handler := err_handler.NewAckingHandler(mockConn)
	result := handler.ProcessErr(mockMsg, err)

	assert.Equal(t, err, result)
}
