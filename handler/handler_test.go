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
	cute.NewTestBuilder().
		Title("Simple test #1").
		Description("Simple get request to '/status'").
		Create().
		RequestBuilder(
			cute.WithURI("http://127.0.0.1:8080/status"),
			cute.WithMethod(http.MethodGet),
		).ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.status", "ok"),
		).ExecuteTest(context.Background(), t)
}
