package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
)

func TestStatusHandler(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected map[string]interface{}
	}{
		{
			name: "Simple Status test",
			url:  "/status",
			expected: map[string]interface{}{
				"status": "ok",
			},
		},
		{
			name: "Simple Healthy Test",
			url:  "/healthy",
			expected: map[string]interface{}{
				"healthy": true,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cute.NewTestBuilder().
				Title(tc.name).
				Create().
				RequestBuilder(
					cute.WithURI("http://127.0.0.1:8080"+tc.url),
					cute.WithMethod(http.MethodGet),
				).ExpectExecuteTimeout(10*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$", tc.expected),
				).ExecuteTest(context.Background(), t)
		})
	}

}
