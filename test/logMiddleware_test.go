package middlewares

import (
	"mini-project/middlewares"

	"testing"

	"github.com/labstack/echo/v4"
)

func TestLogMiddleware(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middlewares.LogMiddleware(tt.args.e)
		})
	}
}
