package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_NotFound(t *testing.T) {
	// Create a temporary directory without config file
	tmpDir := t.TempDir()
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Test with required=false
	cfg, err := LoadConfig(false)
	if err != nil {
		t.Fatalf("LoadConfig(false) should not error when file not found, got: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig(false) should return empty config, not nil")
	}

	// Test with required=true
	_, err = LoadConfig(true)
	if err == nil {
		t.Fatal("LoadConfig(true) should error when file not found")
	}
}

func TestLoadConfig_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Create a valid config file
	configContent := `project:
  name: test-project
  version: 1.0.0
standards:
  path: ./standards
  remote: []
product:
  path: ./product
specs:
  path: ./specs
tasks:
  path: ./tasks
`
	configPath := filepath.Join(tmpDir, ".spekka.yaml")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	cfg, err := LoadConfig(true)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig returned nil config")
	}

	if cfg.Project.Name != "test-project" {
		t.Errorf("Expected project name 'test-project', got '%s'", cfg.Project.Name)
	}
	if cfg.Standards.Path != "./standards" {
		t.Errorf("Expected standards path './standards', got '%s'", cfg.Standards.Path)
	}
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Create an invalid YAML file
	configContent := `project:
  name: test-project
  invalid: yaml: syntax
`
	configPath := filepath.Join(tmpDir, ".spekka.yaml")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	_, err := LoadConfig(true)
	if err == nil {
		t.Fatal("LoadConfig should error on invalid YAML")
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				Standards: StandardsConfig{Path: "./standards"},
				Product:   ProductConfig{Path: "./product"},
				Specs:     SpecsConfig{Path: "./specs"},
				Tasks:     TasksConfig{Path: "./tasks"},
			},
			wantErr: false,
		},
		{
			name: "missing standards path",
			config: &Config{
				Product: ProductConfig{Path: "./product"},
				Specs:   SpecsConfig{Path: "./specs"},
				Tasks:   TasksConfig{Path: "./tasks"},
			},
			wantErr: true,
		},
		{
			name: "missing product path",
			config: &Config{
				Standards: StandardsConfig{Path: "./standards"},
				Specs:     SpecsConfig{Path: "./specs"},
				Tasks:     TasksConfig{Path: "./tasks"},
			},
			wantErr: true,
		},
		{
			name: "missing specs path",
			config: &Config{
				Standards: StandardsConfig{Path: "./standards"},
				Product:   ProductConfig{Path: "./product"},
				Tasks:     TasksConfig{Path: "./tasks"},
			},
			wantErr: true,
		},
		{
			name: "missing tasks path",
			config: &Config{
				Standards: StandardsConfig{Path: "./standards"},
				Product:   ProductConfig{Path: "./product"},
				Specs:     SpecsConfig{Path: "./specs"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


