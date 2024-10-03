package auth

import (
	"strconv"
	"testing"
	"time"
)

func TestSignJWT(t *testing.T) {
	tests := []struct {
		name    string
		userId  int64
		conf    SignConfig
		wantErr bool
	}{
		{
			name:    "simple test",
			userId:  1,
			wantErr: false,
			conf: SignConfig{
				Key:        "secret key",
				ValidUntil: time.Minute,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignJWT(tt.userId, tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignJWT() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got == "" {
				t.Error("SignJWT() got empty string")
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	tests := []struct {
		name    string
		conf    SignConfig
		userId  int64
		token   string
		wantErr bool
	}{
		{
			name:    "success test",
			conf:    SignConfig{Key: "secret test", ValidUntil: time.Minute},
			userId:  1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := SignJWT(tt.userId, tt.conf)
			got, err := ValidateJWT(token, tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got.Subject != strconv.FormatInt(tt.userId, 10) {
				t.Errorf("ValidateJWT() got = %v, want = %v", got.Subject, tt.userId)
			}
		})
	}
}
