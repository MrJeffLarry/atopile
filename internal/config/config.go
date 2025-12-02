package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the global configuration
type Config struct {
	Interactive bool
	Project     *ProjectConfig
}

// ProjectConfig holds project-specific configuration
type ProjectConfig struct {
	Name              string
	Version           string
	Description       string
	OpenLayoutOnBuild bool
	Paths             PathConfig
	Builds            map[string]*BuildConfig
}

// PathConfig holds path-related configuration
type PathConfig struct {
	Root   string
	Logs   string
	Build  string
	Output string
}

// BuildConfig holds build-specific configuration
type BuildConfig struct {
	Name              string
	Entry             string
	Targets           []string
	ExcludeTargets    []string
	KeepPickedParts   bool
	KeepNetNames      bool
	KeepDesignators   bool
	Frozen            bool
	Standalone        bool
	Paths             BuildPathConfig
}

// BuildPathConfig holds build-specific paths
type BuildPathConfig struct {
	Layout     string
	OutputBase string
}

var (
	// config is the global configuration instance
	config = &Config{
		Interactive: true,
		Project:     &ProjectConfig{},
	}
)

// GetConfig returns the global configuration
func GetConfig() *Config {
	return config
}

// ApplyOptions applies command-line options to the configuration
func (c *Config) ApplyOptions(opts *BuildOptions) error {
	if opts == nil {
		return nil
	}

	// Find project root (look for ato.yaml)
	projectRoot, err := FindProjectRoot()
	if err != nil {
		return err
	}

	// TODO: Load ato.yaml configuration file
	c.Project.Paths.Root = projectRoot
	c.Project.Paths.Logs = filepath.Join(projectRoot, ".ato", "logs")
	c.Project.Paths.Build = filepath.Join(projectRoot, ".ato", "build")
	c.Project.Paths.Output = filepath.Join(projectRoot, "build")

	// Apply build-specific options
	if opts.Entry != "" {
		// TODO: parse entry point
	}

	return nil
}

// BuildOptions holds options passed from CLI
type BuildOptions struct {
	Entry           string
	SelectedBuilds  []string
	IncludeTargets  []string
	ExcludeTargets  []string
	Standalone      bool
	Frozen          bool
	KeepPickedParts bool
	KeepNetNames    bool
	KeepDesignators bool
}

// FindProjectRoot searches for the project root directory
// by looking for ato.yaml in the current directory and parent directories
func FindProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		// Check if ato.yaml exists in this directory
		configPath := filepath.Join(dir, "ato.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return dir, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root without finding ato.yaml
			return "", fmt.Errorf("no ato.yaml found in current directory or any parent directory")
		}
		dir = parent
	}
}

// ShouldOpenLayoutOnBuild determines if layout should be opened after build
func (c *Config) ShouldOpenLayoutOnBuild() bool {
	return c.Project.OpenLayoutOnBuild
}
