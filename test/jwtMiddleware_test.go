package middlewares

import (
	"mini-project/middlewares"

	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateToken(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := middlewares.CreateToken(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractTokenUserId(t *testing.T) {
	type args struct {
		e echo.Context
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middlewares.ExtractTokenUserId(tt.args.e); got != tt.want {
				t.Errorf("ExtractTokenUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}
