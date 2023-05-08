package config

import (
	"mini-project/lib/database"
	"mini-project/models"
	"reflect"
	"testing"
)

func TestLoginUsers(t *testing.T) {
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test Login Users",
			args: args{
				user: &models.User{
					Email:    "",
					Password: "",
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "Test Login Users",
			args: args{
				user: &models.User{
					Email:    "rafiandeo26@gmail.com",
					Password: "123456",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := database.LoginUsers(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
