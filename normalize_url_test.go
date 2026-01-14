package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove http scheme",
			inputURL: "http://blog.example.com/path",
			expected: "blog.example.com/path",
		},
		{
			name:     "remove https scheme",
			inputURL: "https://blog.example.com/path",
			expected: "blog.example.com/path",
		},
		{
			name:     "remove training comma",
			inputURL: "https://blog.example.com/path/",
			expected: "blog.example.com/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
