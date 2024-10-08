package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "valid bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer my_token"},
			},
			wantToken: "my_token",
			wantErr:   false,
		},
		{
			name:      "missing authorization header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "invalid authorization header",
			headers: http.Header{
				"Authorization": []string{"Invalid my_token"},
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "malformed bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if token != tt.wantToken {
				t.Errorf("GetBearerToken() = %v, want %v", token, tt.wantToken)
			}
		})
	}
}
