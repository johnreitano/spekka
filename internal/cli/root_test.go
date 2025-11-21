package cli

import (
	"testing"
)

func TestRootCommand(t *testing.T) {
	// Test that root command can be created
	if rootCmd == nil {
		t.Fatal("rootCmd should not be nil")
	}

	if rootCmd.Use != "spekka" {
		t.Errorf("Expected root command use 'spekka', got '%s'", rootCmd.Use)
	}
}

func TestSetupLogging(t *testing.T) {
	tests := []struct {
		name           string
		verbose        bool
		debug          bool
		quiet          bool
		expectedLevel  string
	}{
		{
			name:          "default logging",
			verbose:      false,
			debug:         false,
			quiet:         false,
			expectedLevel: "INFO",
		},
		{
			name:          "verbose logging",
			verbose:      true,
			debug:         false,
			quiet:         false,
			expectedLevel: "INFO",
		},
		{
			name:          "debug logging",
			verbose:      false,
			debug:         true,
			quiet:         false,
			expectedLevel: "DEBUG",
		},
		{
			name:          "quiet logging",
			verbose:      false,
			debug:         false,
			quiet:         true,
			expectedLevel: "ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verbose = tt.verbose
			debug = tt.debug
			quiet = tt.quiet
			setupLogging()
			logger := GetLogger()
			if logger == nil {
				t.Fatal("Logger should not be nil")
			}
		})
	}
}

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	if logger == nil {
		t.Fatal("GetLogger() should not return nil")
	}
}

