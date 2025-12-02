package telemetry

import (
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Errorf("Init() returned error: %v", err)
	}
}

func TestSetEnabled(t *testing.T) {
	SetEnabled(true)
	if !IsEnabled() {
		t.Error("Expected telemetry to be enabled")
	}

	SetEnabled(false)
	if IsEnabled() {
		t.Error("Expected telemetry to be disabled")
	}
}

func TestCapture(t *testing.T) {
	// Should not panic even if disabled
	SetEnabled(false)
	Capture("test_event", nil)
	Capture("test_event", map[string]interface{}{
		"key": "value",
	})

	// Should not panic when enabled
	SetEnabled(true)
	Capture("test_event", nil)
	Capture("test_event", map[string]interface{}{
		"key": "value",
	})
	SetEnabled(false) // Reset
}

func TestShutdown(t *testing.T) {
	// Should not panic
	Shutdown()
}
