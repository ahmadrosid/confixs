package file

import (
	"os"
)

func GetFileList(dirPath string) ([]string, error) {
	var fileList []string

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fileList = append(fileList, entry.Name())
		}
	}

	return fileList, nil
}
