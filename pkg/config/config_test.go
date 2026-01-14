package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandPath(t *testing.T) {
	home, _ := os.UserHomeDir()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "expand home",
			input:    "~/test",
			expected: filepath.Join(home, "test"),
		},
		{
			name:     "absolute path",
			input:    "/absolute/path",
			expected: "/absolute/path",
		},
		{
			name:     "relative path",
			input:    "relative/path",
			expected: "relative/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExpandPath(tt.input)
			if err != nil {
				t.Fatalf("ExpandPath() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("ExpandPath() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetClaudeDir(t *testing.T) {
	home, _ := os.UserHomeDir()
	expected := filepath.Join(home, ".claude")

	result, err := GetClaudeDir()
	if err != nil {
		t.Fatalf("GetClaudeDir() error = %v", err)
	}
	if result != expected {
		t.Errorf("GetClaudeDir() = %v, want %v", result, expected)
	}
}
