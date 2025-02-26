package cmd

import (
	"testing"
)

func TestRot13(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "uryyb"},
		{"HELLO", "URYYB"},
		{"Hello World!", "Uryyb Jbeyq!"},
		{"abcdefghijklmnopqrstuvwxyz", "nopqrstuvwxyzabcdefghijklm"},
		{"1234!@#$", "1234!@#$"},
		{"", ""},
	}

	for _, test := range tests {
		result := rot13(test.input)
		if result != test.expected {
			t.Errorf("rot13(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
