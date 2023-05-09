package config

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
		{
			name: "Test Log Middleware",
			args: args{
				e: echo.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middlewares.LogMiddleware(tt.args.e)
		})
	}
}
