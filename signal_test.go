package warren_test

import(
	"os"
	"syscall"
	"github.com/NGTOne/warren"
	"github.com/NGTOne/warren/conn"

	"github.com/NGTOne/warren/test_mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

type signalGeneratingAction struct{
	signal os.Signal
}

func (act signalGeneratingAction) Process(msg conn.Message) error {
	p, _ := os.FindProcess(os.Getpid())
	return p.Signal(act.signal)
}

func TestDefaultSignalHandler(t *testing.T) {
	tests := []struct{
		name		string
		signal		os.Signal
		shouldShutdown	bool
	}{
		{"SIGINT", syscall.SIGINT, true},
		{"SIGTERM", syscall.SIGTERM, true},
		{"SIGHUP", syscall.SIGHUP, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockMsg := test_mocks.NewMockMessage(mockCtrl)
			mockConn := test_mocks.NewTestConnection(
				mockMsg,
				mockCtrl,
			)
			mockAction := signalGeneratingAction{signal: tt.signal}

			mockMsg.EXPECT().GetHeaderValue("action").Return(
				"foo",
				nil,
			)

			mockConn.EXPECT().AckMsg(mockMsg).Return(nil)

			con := warren.NewConsumer(mockConn)
			con.AddAsyncAction(mockAction, "foo")

			con.Listen()

			if (tt.shouldShutdown) {
				mockConn.EXPECT().Disconnect()
			}
		})
	}
}
