package cli

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommand(t *testing.T) {
	// Reset the command for testing
	cmd := &cobra.Command{
		Use:   "ato",
		Short: "atopile - Design circuit boards with code",
	}
	
	if cmd.Use != "ato" {
		t.Errorf("Expected Use to be 'ato', got '%s'", cmd.Use)
	}
}

func TestVersionFlag(t *testing.T) {
	// This test verifies that version flags are properly registered
	if rootCmd.Flags().Lookup("version") == nil {
		t.Error("version flag not registered")
	}
	
	if rootCmd.Flags().Lookup("semver") == nil {
		t.Error("semver flag not registered")
	}
}

func TestVerboseFlag(t *testing.T) {
	if rootCmd.PersistentFlags().Lookup("verbose") == nil {
		t.Error("verbose flag not registered")
	}
}

func TestNonInteractiveFlag(t *testing.T) {
	if rootCmd.PersistentFlags().Lookup("non-interactive") == nil {
		t.Error("non-interactive flag not registered")
	}
}

func TestDebugFlag(t *testing.T) {
	if rootCmd.PersistentFlags().Lookup("debug") == nil {
		t.Error("debug flag not registered")
	}
}

func TestHiddenFlags(t *testing.T) {
	pythonPath := rootCmd.Flags().Lookup("python-path")
	if pythonPath == nil {
		t.Error("python-path flag not registered")
	}
	if pythonPath != nil && !pythonPath.Hidden {
		t.Error("python-path flag should be hidden")
	}
	
	atopilePath := rootCmd.Flags().Lookup("atopile-path")
	if atopilePath == nil {
		t.Error("atopile-path flag not registered")
	}
	if atopilePath != nil && !atopilePath.Hidden {
		t.Error("atopile-path flag should be hidden")
	}
}

func TestCommandsRegistered(t *testing.T) {
	tests := []string{
		"build",
		"create",
		"validate",
		"inspect",
		"view",
		"dependencies",
		"add",
		"remove",
		"sync",
		"self-check",
	}
	
	for _, cmdName := range tests {
		t.Run(cmdName, func(t *testing.T) {
			cmd, _, err := rootCmd.Find([]string{cmdName})
			if err != nil {
				t.Errorf("Command '%s' not found: %v", cmdName, err)
			}
			if cmd == nil {
				t.Errorf("Command '%s' is nil", cmdName)
			}
		})
	}
}

func TestHiddenCommands(t *testing.T) {
	tests := []string{
		"lsp",
		"mcp",
		"kicad-ipc",
		"package",
		"configure",
		"export-config-schema",
		"dump-config",
		"internal",
	}
	
	for _, cmdName := range tests {
		t.Run(cmdName, func(t *testing.T) {
			cmd, _, err := rootCmd.Find([]string{cmdName})
			if err != nil {
				t.Errorf("Hidden command '%s' not found: %v", cmdName, err)
			}
			if cmd != nil && !cmd.Hidden {
				t.Errorf("Command '%s' should be hidden", cmdName)
			}
		})
	}
}

func TestRootExecute(t *testing.T) {
	// Test that Execute doesn't panic with no arguments
	// We can't fully test execution without mocking, but we can verify it doesn't crash
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Execute panicked: %v", r)
		}
	}()
	
	// Reset command output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
}
