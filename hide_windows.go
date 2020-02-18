// +build windows

package main

import (
	"syscall"
)

func HideFile(path string) (string, error) {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return "", err
	}
	err = syscall.SetFileAttributes(p, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return "", err
	}
	return path, nil
}
