package sig_handler_test

import (
	"os"
	"syscall"
	"github.com/NGTOne/warren/sig_handler"

	"github.com/golang/mock/gomock"
	"testing"
)

func TestShuttingDownHandleSignals(t *testing.T) {
	tests := []struct{
		name		string
		signals		[]os.Signal
		shouldShutdown	bool
	}{
		{"No signals", []os.Signal{}, false},
		{
			"Single unexpected signal",
			[]os.Signal{syscall.SIGHUP},
			false,
		},
		{
			"Multiple different unexpected signals",
			[]os.Signal{syscall.SIGHUP, syscall.SIGALRM},
			false,
		},
		{
			"Multiple of the same unexpected signal",
			[]os.Signal{syscall.SIGHUP, syscall.SIGHUP},
			false,
		},
		{
			"Single expected signal",
			[]os.Signal{syscall.SIGTERM},
			true,
		},
		{
			"Multiple different expected signals",
			[]os.Signal{syscall.SIGTERM, syscall.SIGINT},
			true,
		},
		{
			"Multiple different expected signals (different order)",
			[]os.Signal{syscall.SIGINT, syscall.SIGTERM},
			true,
		},
		{
			"Multiple of the same expected signal",
			[]os.Signal{syscall.SIGTERM, syscall.SIGTERM},
			true,
		},
		{
			"Unexpected signal before expected signal",
			[]os.Signal{syscall.SIGHUP, syscall.SIGTERM},
			true,
		},
		{
			"Unexpected signal after expected signal",
			[]os.Signal{syscall.SIGTERM, syscall.SIGHUP},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockShutdownable := NewMockShutdownable(mockCtrl)

			handler := sig_handler.NewShuttingDownHandler(
				[]os.Signal{
					syscall.SIGTERM,
					syscall.SIGINT,
				},
				mockShutdownable,
			)

			if (tt.shouldShutdown) {
				mockShutdownable.EXPECT().ShutDown()
			}

			handler.HandleSignals(tt.signals)
		})
	}
}
