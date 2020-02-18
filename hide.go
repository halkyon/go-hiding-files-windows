// +build !windows

package main

import (
	"os"
	"path/filepath"
	"strings"
)

func HideFile(path string) (string, error) {
	filename := filepath.Base(path)
	if strings.HasPrefix(filename, ".") {
		return path, nil
	}

	newPath := filepath.Dir(path) + string(os.PathSeparator) + "." + filename
	err := os.Rename(path, newPath)
	if err != nil {
		return "", err
	}

	return newPath, nil
}
