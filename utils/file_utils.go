package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetFileList(dirPath string) ([]string, error) {
	var fileList []string

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fullPath := filepath.Join(dirPath, entry.Name())
			fileList = append(fileList, fullPath)
		}
	}

	return fileList, nil
}

func GetFileContents(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func UpdateFileContents(filePath, content string) error {
	// Check if the file exists
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return errors.New("file does not exist")
	}

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// IsPathSafe checks if the given file path is within the allowed directory
func IsPathSafe(filePath, allowedDir string) bool {
	// Clean the paths to remove any . or .. elements
	cleanFilePath := filepath.Clean(filePath)
	cleanAllowedDir := filepath.Clean(allowedDir)

	// Get the absolute paths
	absFilePath, err := filepath.Abs(cleanFilePath)
	if err != nil {
		return false
	}
	absAllowedDir, err := filepath.Abs(cleanAllowedDir)
	if err != nil {
		return false
	}

	// Use filepath.Rel to get the relative path
	relPath, err := filepath.Rel(absAllowedDir, absFilePath)
	if err != nil {
		return false
	}

	// Check if the relative path starts with ".." or is "."
	// If it does, it means the file is outside the allowed directory
	return !strings.HasPrefix(relPath, "..") && relPath != "."
}
