package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type testInfra struct{
	Ctrl *gomock.Controller
	Class *MockClassifier
	Resp *MockReponder
}

func (t *testInfra) Finish() {
	t.Ctrl.Finish()
}

func newTestInfra(t *testing.T) *testInfra {
	ctrl := gomock.NewController(t)
	class := NewMockClassifier(ctrl)
	resp := NewMockReponder(ctrl)
	return &testInfra{
		Ctrl:  ctrl,
		Class: class,
		Resp:  resp,
	}
}

func TestService_ListenAndRespond(t *testing.T) {
	tests := []struct {
		desc string
		input string
		class string
		response string
		expectedResponse string
	}{
		{
			desc: "should succeed",
			input: "test",
			class: "testClass",
			response: "testResp",
			expectedResponse: "testResp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			testInfra := newTestInfra(t)
			defer testInfra.Finish()

			testInfra.Class.EXPECT().Classify(tt.input).Return(tt.class)
			testInfra.Resp.EXPECT().Respond(tt.input, tt.class).Return(tt.response)

			service := NewService(testInfra.Class, testInfra.Resp)
			serveResp := service.ListenAndRespond(tt.input)

			require.Equal(t,tt.expectedResponse,serveResp)
		})
	}
}
