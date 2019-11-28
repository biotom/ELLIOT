package responder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponder_Respond(t *testing.T) {
	tests := []struct {
		desc string
		input string
		classification string
	}{
		{
			desc: "should successfully return string",
			input: "hi",
			classification: "hi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			responder := NewResponder()
			require.NotEmpty(t, responder.Respond(tt.input, tt.classification))
		})
	}
}
