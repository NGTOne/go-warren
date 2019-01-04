package sig_handler_test

import (
	"os"
	"syscall"
	"github.com/NGTOne/warren/sig_handler"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandleSignals(t *testing.T) {
	handler := sig_handler.NewPanickingHandler(map[os.Signal]string{
		syscall.SIGTERM: "Caught SIGTERM",
		syscall.SIGINT:  "Caught SIGINT",
	})

	tests := []struct{
		name		string
		signals		[]os.Signal
		shouldPanic	bool
		expectedMsg	string
	}{
		{"No signals", []os.Signal{}, false, ""},
		{
			"Single unexpected signal",
			[]os.Signal{syscall.SIGHUP},
			false,
			"",
		},
		{
			"Multiple different unexpected signals",
			[]os.Signal{syscall.SIGHUP, syscall.SIGALRM},
			false,
			"",
		},
		{
			"Multiple of the same unexpected signal",
			[]os.Signal{syscall.SIGHUP, syscall.SIGHUP},
			false,
			"",
		},
		{
			"Single expected signal",
			[]os.Signal{syscall.SIGTERM},
			true,
			"Caught SIGTERM",
		},
		{
			"Multiple different expected signals",
			[]os.Signal{syscall.SIGTERM, syscall.SIGINT},
			true,
			"Caught SIGTERM",
		},
		{
			"Multiple different expected signals (different order)",
			[]os.Signal{syscall.SIGINT, syscall.SIGTERM},
			true,
			"Caught SIGINT",
		},
		{
			"Multiple of the same expected signal",
			[]os.Signal{syscall.SIGTERM, syscall.SIGTERM},
			true,
			"Caught SIGTERM",
		},
		{
			"Unexpected signal before expected signal",
			[]os.Signal{syscall.SIGHUP, syscall.SIGTERM},
			true,
			"Caught SIGTERM",
		},
		{
			"Unexpected signal after expected signal",
			[]os.Signal{syscall.SIGTERM, syscall.SIGHUP},
			true,
			"Caught SIGTERM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunc := func() {
				handler.HandleSignals(tt.signals)
			}

			if (tt.shouldPanic) {
				assert.PanicsWithValue(t, tt.expectedMsg, testFunc)
			} else {
				assert.NotPanics(t, testFunc)
			}
		})
	}
}
