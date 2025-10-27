package version

import "fmt"

const (
	Major = 0
	Minor = 1
	Patch = 0
)

// String returns the semantic version string.
func String() string {
	return fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
}
