package auth

import (
	"errors"
	"testing"
)

type testCase struct {
	name     string
	input    map[string][]string
	expected string
	err      error
}

func TestGetAPIKey(t *testing.T) {
	testCases := []testCase{
		{
			name:     "getting api key",
			input:    map[string][]string{"Authorization": {"ApiKey 123"}},
			expected: "123",
			err:      nil,
		},
		{
			name:     "getting an error with message, because there are no headers",
			input:    map[string][]string{"": {""}},
			expected: "",
			err:      errors.New("no authorization header included"),
		},
		{
			name:     "getting an error with message, because there authorization is not apikey",
			input:    map[string][]string{"Authorization": {"Bearer 123"}},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
		{
			name:     "getting an error with message, because there authorization format is incorrect",
			input:    map[string][]string{"Authorization": {"ApiKey123"}},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.input)

			if apiKey != tc.expected {
				t.Fatalf("expected: %v, got %v", tc.expected, apiKey)
			}

			if err == nil && tc.err != nil {
				t.Fatalf("expected error: %v, got nil", tc.err)
			}

			if err != nil && tc.err == nil {
				t.Fatalf("expected no error, got: %v", err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("expected error: %v, got %v", tc.err.Error(), err.Error())
			}
		})
	}
}
