// Package config provides configuration management for Spekka.
// It handles loading and parsing .spekka.yaml configuration files using Viper,
// with support for configuration file discovery and validation.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config represents the Spekka configuration structure.
type Config struct {
	Project   ProjectConfig   `mapstructure:"project"`
	Standards StandardsConfig `mapstructure:"standards"`
	Product   ProductConfig   `mapstructure:"product"`
	Specs     SpecsConfig     `mapstructure:"specs"`
	Tasks     TasksConfig     `mapstructure:"tasks"`
}

// ProjectConfig contains project-level configuration.
type ProjectConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

// StandardsConfig contains standards layer configuration.
type StandardsConfig struct {
	Path   string   `mapstructure:"path"`
	Remote []string `mapstructure:"remote"`
}

// ProductConfig contains product vision configuration.
type ProductConfig struct {
	Path string `mapstructure:"path"`
}

// SpecsConfig contains specifications configuration.
type SpecsConfig struct {
	Path string `mapstructure:"path"`
}

// TasksConfig contains tasks configuration.
type TasksConfig struct {
	Path string `mapstructure:"path"`
}

// LoadConfig loads configuration from .spekka.yaml file.
// It searches for the config file in the current directory and parent directories.
// Returns an error if the config file is required but not found, or if the file
// contains invalid YAML.
func LoadConfig(required bool) (*Config, error) {
	v := viper.New()
	v.SetConfigName(".spekka")
	v.SetConfigType("yaml")

	// Search for config file starting from current directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	// Search current directory and parent directories
	searchPath := wd
	for {
		v.AddConfigPath(searchPath)
		if searchPath == filepath.Dir(searchPath) {
			break // Reached filesystem root
		}
		searchPath = filepath.Dir(searchPath)
	}

	// Also check home directory
	if home, err := os.UserHomeDir(); err == nil {
		v.AddConfigPath(home)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if required {
				return nil, fmt.Errorf("configuration file .spekka.yaml not found: %w", err)
			}
			// Return empty config if not required
			return &Config{}, nil
		}
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}

	return &cfg, nil
}

// Validate checks that required configuration fields are present.
// Returns an error listing all missing required fields.
func (c *Config) Validate() error {
	var missing []string

	if c.Standards.Path == "" {
		missing = append(missing, "standards.path")
	}
	if c.Product.Path == "" {
		missing = append(missing, "product.path")
	}
	if c.Specs.Path == "" {
		missing = append(missing, "specs.path")
	}
	if c.Tasks.Path == "" {
		missing = append(missing, "tasks.path")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required configuration fields: %v", missing)
	}

	return nil
}

