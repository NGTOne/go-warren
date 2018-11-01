package err_handler_test

import (
	"github.com/NGTOne/warren/err_handler"
	"github.com/NGTOne/warren/test_mocks"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNackingHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockConn.EXPECT().NackMsg(mockMsg).Return(nil)

	err := errors.New("Something went wrong!")

	handler := err_handler.NewNackingHandler(mockConn)
	result := handler.ProcessErr(mockMsg, err)

	assert.Equal(t, err, result)
}
