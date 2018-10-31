package err_handler_test

import(
	"github.com/NGTOne/warren/err_handler"
	"github.com/NGTOne/warren/test_mocks"

	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestCompositionalHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMsg := test_mocks.NewMockMessage(mockCtrl)
	mockConn := test_mocks.NewTestConnection(mockMsg, mockCtrl)

	mockConn.EXPECT().AckMessage(mockMsg).Return(nil)
	mockConn.EXPECT().NackMessage(mockMsg).Return(nil)

	err := errors.New("Something went wrong!")

	handler := err_handler.NewCompositionalHandler(
		[]err_handler.ErrorHandler{
			// Obviously this setup is ridiculous in practice, but
			// we just need two error handlers here to test things
			err_handler.NewAckingHandler(mockConn),
			err_handler.NewNackingHandler(mockConn),
		},
	)

	result := handler.ProcessError(mockMsg, err)

	assert.Equal(t, err, result)
}
