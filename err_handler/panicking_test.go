package err_handler_test

import (
	"github.com/NGTOne/warren/err_handler"
	"github.com/NGTOne/warren/test_mocks"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanickingHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMsg := test_mocks.NewMockMessage(mockCtrl)

	err := errors.New("Something went wrong!")
	handler := err_handler.PanickingHandler{}

	assert.Panics(t, func() { handler.ProcessErr(mockMsg, err) })
}
