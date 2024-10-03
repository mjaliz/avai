package auth

import (
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
