package auth_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		sep   string
		want  string
		err   bool
	}{
		"valid header": {input: http.Header{
			"Authorization": []string{"ApiKey est"},
		}, want: "test"},
		"empty header": {input: http.Header{}, err: true},
		"invalid auth scheme": {input: http.Header{
			"Authorization": []string{"Broken test"},
		}, err: true},
		"no auth token, but scheme is valid": {input: http.Header{
			"Authorization": []string{"ApiKey"},
		}, err: true},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := auth.GetAPIKey(tc.input)
			if err != nil && tc.err != true {
				t.Fatalf("expected to not error, got: %v", err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
