package version

import (
	"fmt"
	"runtime/debug"
)

// Version information - these should be set during build
var (
	Version   = "dev"
	GitCommit = ""
	BuildDate = ""
)

// GetVersion returns the current version string
func GetVersion() string {
	if Version != "dev" {
		return Version
	}

	// Try to get version from build info
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" && info.Main.Version != "" {
			return info.Main.Version
		}
	}

	return "dev"
}

// GetVersionInfo returns detailed version information
func GetVersionInfo() string {
	v := GetVersion()
	if GitCommit != "" {
		v += fmt.Sprintf(" (commit: %s)", GitCommit)
	}
	if BuildDate != "" {
		v += fmt.Sprintf(" (built: %s)", BuildDate)
	}
	return v
}

// Parse parses a version string into semver format
// This is a simplified version - for full semver support, use a library
func Parse(versionStr string) string {
	// For now, just return the version as-is
	// TODO: implement proper semver parsing
	return versionStr
}

// CheckForUpdate checks if a new version is available
func CheckForUpdate() {
	// TODO: implement update checking
	// This would check against PyPI or a similar registry
}
