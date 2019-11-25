package classifier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClassifier_Classify(t *testing.T) {
	tests := []struct {
		desc string
		input string
		expectedClassification string
	}{
		{
			desc: "should succeed",
			input: "whatever",
			expectedClassification:"filler",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c := NewClassifier()
			result := c.Classify(tt.input)
			require.Equal(t, tt.expectedClassification, result)
		})
	}
}
