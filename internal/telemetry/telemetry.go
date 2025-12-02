package telemetry

import (
	"fmt"
)

// Config holds telemetry configuration
type Config struct {
	Enabled bool
	APIKey  string
}

var (
	config = &Config{
		Enabled: false, // Disabled by default
	}
)

// Init initializes the telemetry system
func Init() error {
	// TODO: Initialize telemetry backend (e.g., PostHog)
	// For now, this is a stub
	return nil
}

// Capture sends a telemetry event
func Capture(eventName string, properties map[string]interface{}) {
	if !config.Enabled {
		return
	}

	// TODO: Send event to telemetry backend
	// For now, just log in debug mode
	if properties != nil {
		fmt.Printf("Telemetry: %s - %v\n", eventName, properties)
	}
}

// SetEnabled enables or disables telemetry
func SetEnabled(enabled bool) {
	config.Enabled = enabled
}

// IsEnabled returns whether telemetry is enabled
func IsEnabled() bool {
	return config.Enabled
}

// Shutdown gracefully shuts down the telemetry system
func Shutdown() {
	// TODO: Flush any pending events
}
