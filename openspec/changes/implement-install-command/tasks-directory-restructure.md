# Tasks: Directory Structure Modification

## 1. Implementation
- [ ] 1.1 Update spekkaDirectories variable to include .spekka parent path
- [ ] 1.2 Modify createDirectoryStructure() to create .spekka/ parent directory first
- [ ] 1.3 Update directory creation logic to create subdirectories under .spekka/
- [ ] 1.4 Modify getConfigTemplate() to use .spekka/ prefixed paths
- [ ] 1.5 Update createTemplateFiles() to write files to .spekka/ subdirectories
- [ ] 1.6 Update displaySuccessMessage() to show new directory structure
- [ ] 1.7 Modify runDryRun() to display new nested structure

## 2. Configuration Updates
- [ ] 2.1 Update default configuration template paths to use .spekka/ prefix
- [ ] 2.2 Ensure configuration validation works with new paths
- [ ] 2.3 Update configuration comments to reflect new structure
- [ ] 2.4 Verify YAML structure remains valid with new paths

## 3. Testing Updates
- [ ] 3.1 Update unit tests to expect .spekka/ parent directory
- [ ] 3.2 Modify integration tests to validate new directory structure
- [ ] 3.3 Update test assertions to check for .spekka/ subdirectories
- [ ] 3.4 Verify template file creation in correct nested locations
- [ ] 3.5 Test configuration file paths point to correct directories
- [ ] 3.6 Update dry-run tests to expect new structure display
- [ ] 3.7 Test force reinstallation with existing .spekka/ directory

## 4. Backward Compatibility
- [ ] 4.1 Consider migration strategy for existing installations
- [ ] 4.2 Add detection for old directory structure
- [ ] 4.3 Provide migration guidance in documentation
- [ ] 4.4 Update error messages to reference new structure

## 5. Documentation Updates
- [ ] 5.1 Update template README files to reference new paths
- [ ] 5.2 Update help text and command descriptions
- [ ] 5.3 Modify success messages to show .spekka/ structure
- [ ] 5.4 Update any hardcoded path references in templates

## 6. Validation
- [ ] 6.1 Run all existing tests to ensure no regressions
- [ ] 6.2 Verify directory permissions are correct for .spekka/
- [ ] 6.3 Test installation in various scenarios (empty dir, existing files)
- [ ] 6.4 Validate configuration file parsing with new paths
- [ ] 6.5 Manual smoke test of complete installation flow
- [ ] 6.6 Verify .gitkeep files are created in correct locations
