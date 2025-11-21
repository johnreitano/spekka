package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunInstall(t *testing.T) {
	tests := []struct {
		name           string
		setupFunc      func(t *testing.T, tempDir string)
		flags          map[string]bool
		expectError    bool
		errorContains  string
		validateFunc   func(t *testing.T, tempDir string)
	}{
		{
			name: "fresh installation success",
			setupFunc: func(t *testing.T, tempDir string) {
				// No setup needed for fresh installation
			},
			flags:       map[string]bool{},
			expectError: false,
			validateFunc: func(t *testing.T, tempDir string) {
				// Check directories were created
				for _, dir := range spekkaDirectories {
					dirPath := filepath.Join(tempDir, dir)
					assert.DirExists(t, dirPath, "Directory %s should exist", dir)
					
					// Check README.md exists
					readmePath := filepath.Join(dirPath, "README.md")
					assert.FileExists(t, readmePath, "README.md should exist in %s", dir)
					
					// Check .gitkeep exists
					gitkeepPath := filepath.Join(dirPath, ".gitkeep")
					assert.FileExists(t, gitkeepPath, ".gitkeep should exist in %s", dir)
				}
				
				// Check config file was created
				configPath := filepath.Join(tempDir, ".spekka/config.yaml")
				assert.FileExists(t, configPath, ".spekka/config.yaml should exist")
				
				// Verify config content
				content, err := os.ReadFile(configPath)
				require.NoError(t, err)
				assert.Contains(t, string(content), "project:")
				assert.Contains(t, string(content), "standards:")
			},
		},
		{
			name: "existing installation without force",
			setupFunc: func(t *testing.T, tempDir string) {
				// Create existing .spekka/config.yaml
				err := os.MkdirAll(filepath.Join(tempDir, ".spekka"), 0755)
				require.NoError(t, err)
				configPath := filepath.Join(tempDir, ".spekka/config.yaml")
				err = os.WriteFile(configPath, []byte("existing config"), 0644)
				require.NoError(t, err)
			},
			flags:         map[string]bool{},
			expectError:   true,
			errorContains: "already installed",
		},
		{
			name: "force reinstallation",
			setupFunc: func(t *testing.T, tempDir string) {
				// Create existing .spekka/config.yaml
				err := os.MkdirAll(filepath.Join(tempDir, ".spekka"), 0755)
				require.NoError(t, err)
				configPath := filepath.Join(tempDir, ".spekka/config.yaml")
				err = os.WriteFile(configPath, []byte("existing config"), 0644)
				require.NoError(t, err)
			},
			flags:       map[string]bool{"force": true},
			expectError: false,
			validateFunc: func(t *testing.T, tempDir string) {
				// Check backup was created
				backupFiles, err := filepath.Glob(filepath.Join(tempDir, ".spekka/config.yaml.backup-*"))
				require.NoError(t, err)
				assert.Len(t, backupFiles, 1, "Backup file should be created")
				
				// Check new config exists
				configPath := filepath.Join(tempDir, ".spekka/config.yaml")
				assert.FileExists(t, configPath)
				
				// Verify it's the new config, not the old one
				content, err := os.ReadFile(configPath)
				require.NoError(t, err)
				assert.NotEqual(t, "existing config", string(content))
			},
		},
		{
			name: "dry run mode",
			setupFunc: func(t *testing.T, tempDir string) {
				// No setup needed
			},
			flags:       map[string]bool{"dry-run": true},
			expectError: false,
			validateFunc: func(t *testing.T, tempDir string) {
				// Check that nothing was actually created
				configPath := filepath.Join(tempDir, ".spekka/config.yaml")
				assert.NoFileExists(t, configPath, "Config file should not exist in dry run")
				
				for _, dir := range spekkaDirectories {
					dirPath := filepath.Join(tempDir, dir)
					assert.NoDirExists(t, dirPath, "Directory %s should not exist in dry run", dir)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tempDir := t.TempDir()
			
			// Change to temp directory
			originalDir, err := os.Getwd()
			require.NoError(t, err)
			defer func() {
				err := os.Chdir(originalDir)
				require.NoError(t, err)
			}()
			err = os.Chdir(tempDir)
			require.NoError(t, err)
			
			// Setup test environment
			if tt.setupFunc != nil {
				tt.setupFunc(t, tempDir)
			}
			
			// Create command with flags
			cmd := &cobra.Command{}
			for flag, value := range tt.flags {
				cmd.Flags().Bool(flag, value, "")
				if value {
					err := cmd.Flags().Set(flag, "true")
					require.NoError(t, err)
				}
			}
			
			// Run the install command
			err = runInstall(cmd, []string{})
			
			// Check error expectation
			if tt.expectError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
			}
			
			// Run validation function
			if tt.validateFunc != nil {
				tt.validateFunc(t, tempDir)
			}
		})
	}
}

func TestConfigFileExists(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T, tempDir string)
		expected bool
	}{
		{
			name: "config file exists",
			setup: func(t *testing.T, tempDir string) {
				err := os.MkdirAll(filepath.Join(tempDir, ".spekka"), 0755)
				require.NoError(t, err)
				err = os.WriteFile(filepath.Join(tempDir, ".spekka/config.yaml"), []byte("test"), 0644)
				require.NoError(t, err)
			},
			expected: true,
		},
		{
			name: "config file does not exist",
			setup: func(t *testing.T, tempDir string) {
				// No setup - file doesn't exist
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			
			// Change to temp directory
			originalDir, err := os.Getwd()
			require.NoError(t, err)
			defer func() {
				err := os.Chdir(originalDir)
				require.NoError(t, err)
			}()
			err = os.Chdir(tempDir)
			require.NoError(t, err)
			
			if tt.setup != nil {
				tt.setup(t, tempDir)
			}
			
			exists, err := configFileExists()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, exists)
		})
	}
}

func TestCreateDirectoryStructure(t *testing.T) {
	tempDir := t.TempDir()
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	
	err = createDirectoryStructure()
	require.NoError(t, err)
	
	// Verify all directories were created
	for _, dir := range spekkaDirectories {
		dirPath := filepath.Join(tempDir, dir)
		assert.DirExists(t, dirPath, "Directory %s should exist", dir)
		
		// Check permissions
		info, err := os.Stat(dirPath)
		require.NoError(t, err)
		assert.True(t, info.IsDir())
		assert.Equal(t, os.FileMode(0755), info.Mode().Perm())
	}
}

func TestGenerateConfigFile(t *testing.T) {
	tempDir := t.TempDir()
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	
	// Create .spekka directory first
	err = os.MkdirAll(".spekka", 0755)
	require.NoError(t, err)
	
	err = generateConfigFile()
	require.NoError(t, err)
	
	// Verify config file was created
	configPath := filepath.Join(tempDir, ".spekka/config.yaml")
	assert.FileExists(t, configPath)
	
	// Verify content
	content, err := os.ReadFile(configPath)
	require.NoError(t, err)
	
	configStr := string(content)
	assert.Contains(t, configStr, "project:")
	assert.Contains(t, configStr, "standards:")
	assert.Contains(t, configStr, "product:")
	assert.Contains(t, configStr, "specs:")
	assert.Contains(t, configStr, "tasks:")
	assert.Contains(t, configStr, "path: \"standards\"")
	assert.Contains(t, configStr, "path: \"product\"")
}

func TestCreateTemplateFiles(t *testing.T) {
	tempDir := t.TempDir()
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	
	// Create directories first
	err = createDirectoryStructure()
	require.NoError(t, err)
	
	err = createTemplateFiles()
	require.NoError(t, err)
	
	// Verify template files were created
	for _, dir := range spekkaDirectories {
		readmePath := filepath.Join(tempDir, dir, "README.md")
		assert.FileExists(t, readmePath, "README.md should exist in %s", dir)
		
		// Verify content is not empty
		content, err := os.ReadFile(readmePath)
		require.NoError(t, err)
		assert.NotEmpty(t, content, "README.md in %s should not be empty", dir)
		
		// Verify .gitkeep files
		gitkeepPath := filepath.Join(tempDir, dir, ".gitkeep")
		assert.FileExists(t, gitkeepPath, ".gitkeep should exist in %s", dir)
	}
}

func TestTemplateContent(t *testing.T) {
	tests := []struct {
		name     string
		template func() string
		contains []string
	}{
		{
			name:     "standards template",
			template: getStandardsTemplate,
			contains: []string{"Standards Directory", "coding standards", "3-layer context", "Git URL"},
		},
		{
			name:     "product template",
			template: getProductTemplate,
			contains: []string{"Product Directory", "product vision", "mission.md", "personas"},
		},
		{
			name:     "specs template",
			template: getSpecsTemplate,
			contains: []string{"Specifications Directory", "requirements", "SHALL provide", "Scenario:"},
		},
		{
			name:     "tasks template",
			template: getTasksTemplate,
			contains: []string{"Tasks Directory", "YAML frontmatter", "effort:", "status:"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := tt.template()
			assert.NotEmpty(t, content)
			
			for _, expected := range tt.contains {
				assert.Contains(t, content, expected, "Template should contain '%s'", expected)
			}
		})
	}
}

func TestBackupExistingConfig(t *testing.T) {
	tempDir := t.TempDir()
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	
	// Create existing config
	originalContent := "original config content"
	err = os.MkdirAll(".spekka", 0755)
	require.NoError(t, err)
	err = os.WriteFile(".spekka/config.yaml", []byte(originalContent), 0644)
	require.NoError(t, err)
	
	err = backupExistingConfig()
	require.NoError(t, err)
	
	// Verify backup was created
	backupFiles, err := filepath.Glob(".spekka/config.yaml.backup-*")
	require.NoError(t, err)
	assert.Len(t, backupFiles, 1, "Should create exactly one backup file")
	
	// Verify backup content
	backupContent, err := os.ReadFile(backupFiles[0])
	require.NoError(t, err)
	assert.Equal(t, originalContent, string(backupContent))
	
	// Verify original file is gone
	_, err = os.Stat(".spekka/config.yaml")
	assert.True(t, os.IsNotExist(err), "Original config should be moved")
}

func TestRunDryRun(t *testing.T) {
	tempDir := t.TempDir()
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	
	err = runDryRun()
	require.NoError(t, err)
	
	// Verify nothing was actually created
	_, err = os.Stat(".spekka/config.yaml")
	assert.True(t, os.IsNotExist(err), "Config file should not exist after dry run")
	
	for _, dir := range spekkaDirectories {
		_, err = os.Stat(dir)
		assert.True(t, os.IsNotExist(err), "Directory %s should not exist after dry run", dir)
	}
}

func TestGetConfigTemplate(t *testing.T) {
	config := getConfigTemplate()
	
	assert.NotEmpty(t, config)
	assert.Contains(t, config, "project:")
	assert.Contains(t, config, "name: \"my-project\"")
	assert.Contains(t, config, "standards:")
	assert.Contains(t, config, "path: \"standards\"")
	assert.Contains(t, config, "product:")
	assert.Contains(t, config, "specs:")
	assert.Contains(t, config, "tasks:")
	
	// Verify it's valid YAML structure (basic check)
	lines := strings.Split(config, "\n")
	var hasProjectSection, hasStandardsSection bool
	for _, line := range lines {
		if strings.HasPrefix(line, "project:") {
			hasProjectSection = true
		}
		if strings.HasPrefix(line, "standards:") {
			hasStandardsSection = true
		}
	}
	assert.True(t, hasProjectSection, "Config should have project section")
	assert.True(t, hasStandardsSection, "Config should have standards section")
}
