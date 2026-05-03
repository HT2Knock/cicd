package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "Valid token provided",
			headers: http.Header{
				"Authorization": []string{"ApiKey validToken"},
			},
			want:    "validToken",
			wantErr: false,
		},

		{
			name: "Invalid token provided",
			headers: http.Header{
				"Authorization": []string{"Bearer: test"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "No auth header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}

				return
			}

			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
