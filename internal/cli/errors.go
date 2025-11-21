package cli

import (
	"fmt"
	"os"
)

// FormatError formats an error message for user display.
// It extracts context from wrapped errors and presents them clearly.
func FormatError(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("Error: %v", err)
}

// HandleError writes an error to stderr and exits with the appropriate exit code.
func HandleError(err error, exitCode int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", FormatError(err))
		os.Exit(exitCode)
	}
}

// ExitWithError writes an error message to stderr and exits with code 1.
func ExitWithError(err error) {
	HandleError(err, 1)
}

// ExitWithErrorCode writes an error message to stderr and exits with the specified code.
func ExitWithErrorCode(err error, code int) {
	HandleError(err, code)
}


