// Package cli provides the command-line interface for Spekka.
// It implements the root command and all subcommands using the Cobra framework,
// handles logging configuration, and provides error handling utilities.
package cli

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var (
	version  = "0.0.1"
	verbose  bool
	debug    bool
	quiet    bool
	logLevel slog.Level
	logger   *slog.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spekka",
	Short: "Spekka - Spec-driven development workflow tool",
	Long: `Spekka is a CLI-based development workflow tool that helps software
engineering teams build better products through spec-driven development.

It provides a structured, AI-assisted approach to planning, specification
writing, and implementation that maximizes the value of LLM context windows.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check for version flag (--version)
		// Note: -v is handled as verbose (persistent flag) for subcommands
		// For root command, use --version to show version
		if v, _ := cmd.Flags().GetBool("version"); v {
			cmd.Println("spekka version", version)
			return nil
		}
		// Otherwise show help
		return cmd.Help()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Set up logging based on flags
		setupLogging()
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress non-error output")

	// Version flag (not persistent, only on root)
	// Use --version for long form, handle -v specially in RunE
	rootCmd.Flags().Bool("version", false, "Display version information")
	rootCmd.Flags().Lookup("version").NoOptDefVal = "true"
}

// setupLogging configures the logger based on command flags.
func setupLogging() {
	// Determine log level based on flags
	if quiet {
		logLevel = slog.LevelError
	} else if debug {
		logLevel = slog.LevelDebug
	} else if verbose {
		logLevel = slog.LevelInfo
	} else {
		logLevel = slog.LevelInfo
	}

	// Create logger with appropriate level
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger = slog.New(handler)
}

// GetLogger returns the configured logger instance.
func GetLogger() *slog.Logger {
	if logger == nil {
		// Default logger if not yet initialized
		setupLogging()
	}
	return logger
}
