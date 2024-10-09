package main

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			wantKey: "1234567890",
		},
		{
			name:    "missing header",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name: "invalid header format",
			headers: http.Header{
				"Authorization": []string{"Invalid 1234567890"},
			},
			wantErr: true,
		},
		{
			name: "missing ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"1234567890"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if key != tt.wantKey {
				t.Errorf("GetAPIKey() = %v, want %v", key, tt.wantKey)
			}
		})
	}
}
