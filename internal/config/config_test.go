package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := GetConfig()
	if cfg == nil {
		t.Error("GetConfig returned nil")
	}

	if !cfg.Interactive {
		t.Error("Expected Interactive to be true by default")
	}
}

func TestBuildOptions(t *testing.T) {
	opts := &BuildOptions{
		Entry:          "test.ato:Module",
		SelectedBuilds: []string{"build1"},
		IncludeTargets: []string{"target1"},
	}

	if opts.Entry != "test.ato:Module" {
		t.Errorf("Expected Entry to be 'test.ato:Module', got '%s'", opts.Entry)
	}

	if len(opts.SelectedBuilds) != 1 {
		t.Errorf("Expected 1 selected build, got %d", len(opts.SelectedBuilds))
	}

	if len(opts.IncludeTargets) != 1 {
		t.Errorf("Expected 1 include target, got %d", len(opts.IncludeTargets))
	}
}

func TestProjectConfig(t *testing.T) {
	cfg := &ProjectConfig{
		Name:              "test-project",
		Version:           "1.0.0",
		OpenLayoutOnBuild: true,
	}

	if cfg.Name != "test-project" {
		t.Errorf("Expected Name to be 'test-project', got '%s'", cfg.Name)
	}

	if cfg.Version != "1.0.0" {
		t.Errorf("Expected Version to be '1.0.0', got '%s'", cfg.Version)
	}

	if !cfg.OpenLayoutOnBuild {
		t.Error("Expected OpenLayoutOnBuild to be true")
	}
}

func TestBuildConfig(t *testing.T) {
	cfg := &BuildConfig{
		Name:    "default",
		Entry:   "main.ato:Main",
		Targets: []string{"target1", "target2"},
		Frozen:  true,
	}

	if cfg.Name != "default" {
		t.Errorf("Expected Name to be 'default', got '%s'", cfg.Name)
	}

	if len(cfg.Targets) != 2 {
		t.Errorf("Expected 2 targets, got %d", len(cfg.Targets))
	}

	if !cfg.Frozen {
		t.Error("Expected Frozen to be true")
	}
}

func TestShouldOpenLayoutOnBuild(t *testing.T) {
	cfg := &Config{
		Project: &ProjectConfig{
			OpenLayoutOnBuild: true,
		},
	}

	if !cfg.ShouldOpenLayoutOnBuild() {
		t.Error("Expected ShouldOpenLayoutOnBuild to return true")
	}

	cfg.Project.OpenLayoutOnBuild = false
	if cfg.ShouldOpenLayoutOnBuild() {
		t.Error("Expected ShouldOpenLayoutOnBuild to return false")
	}
}
