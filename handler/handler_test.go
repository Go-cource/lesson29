package handler

import (
	"testing"
)

func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		// TODO: Add test cases.
		{
			name: "Simple test #1",
			want: want{
				code:        200,
				response:    `{"status": "ok"}`,
				contentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StatusHandler(tt.args.w, tt.args.r)
		})
	}
}
