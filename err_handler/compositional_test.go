package err_handler_test

import (
	"github.com/NGTOne/warren/err_handler"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompositionalHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := NewMockMessage(mockCtrl)
	mockConn := NewMockConnection(mockCtrl)

	mockConn.EXPECT().AckMsg(mockMsg).Return(nil)
	mockConn.EXPECT().NackMsg(mockMsg).Return(nil)

	err := errors.New("Something went wrong!")

	handler := err_handler.NewCompositionalHandler(
		[]err_handler.ErrHandler{
			// Obviously this setup is ridiculous in practice, but
			// we just need two error handlers here to test things
			err_handler.NewAckingHandler(mockConn),
			err_handler.NewNackingHandler(mockConn),
		},
	)

	result := handler.ProcessErr(mockMsg, err)

	assert.Equal(t, err, result)
}
