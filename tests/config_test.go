package config

import (
	"mini-project/config"

	"testing"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Init DB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitDB()
		})
	}
}

func TestInitMigrate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Init Migrate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitMigrate()
		})
	}
}
