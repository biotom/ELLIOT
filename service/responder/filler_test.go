package responder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFillerResponse(t *testing.T) {
	tests := []struct {
		desc string
		callTimes int
	}{
		{
			desc: "should successfully return string",
			callTimes: 1,
		},
		{
			desc: "should succeed when called more times than there are default filler phrases",
			callTimes: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			i := 0
			for i < tt.callTimes{
				fillerResponse := fillerResponder()
				require.NotEmpty(t,  fillerResponse())
				i += 1
			}

		})
	}
}