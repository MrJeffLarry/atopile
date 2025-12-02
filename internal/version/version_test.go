package version

import (
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v == "" {
		t.Error("GetVersion returned empty string")
	}
}

func TestGetVersionInfo(t *testing.T) {
	// Set test values
	oldVersion := Version
	oldCommit := GitCommit
	oldBuildDate := BuildDate
	
	Version = "1.0.0"
	GitCommit = "abc123"
	BuildDate = "2024-01-01"
	
	defer func() {
		Version = oldVersion
		GitCommit = oldCommit
		BuildDate = oldBuildDate
	}()
	
	info := GetVersionInfo()
	
	if !strings.Contains(info, "1.0.0") {
		t.Errorf("Expected version info to contain '1.0.0', got: %s", info)
	}
	
	if !strings.Contains(info, "abc123") {
		t.Errorf("Expected version info to contain commit hash, got: %s", info)
	}
	
	if !strings.Contains(info, "2024-01-01") {
		t.Errorf("Expected version info to contain build date, got: %s", info)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1.0.0", "1.0.0"},
		{"v1.2.3", "v1.2.3"},
		{"2.0.0-beta.1", "2.0.0-beta.1"},
	}
	
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Parse(tt.input)
			if result != tt.expected {
				t.Errorf("Parse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCheckForUpdate(t *testing.T) {
	// This should not panic
	CheckForUpdate()
}
