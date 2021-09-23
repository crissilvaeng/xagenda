package support

import "fmt"

// MissingConfigError is an error type for missing config.
// Occurs when a config lacks of default value and has not been set.
type MissingConfigError struct {
	Field string
}

// Error returns a string representation of the error.
func (e *MissingConfigError) Error() string {
	return fmt.Sprintf("missing required configL %v", e.Field)
}
