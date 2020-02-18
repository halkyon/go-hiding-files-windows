package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	dir := os.TempDir()
	path := filepath.Join(dir, "testfile.txt")
	err := ioutil.WriteFile(path, []byte("hello\n"), 0600)
	if err != nil {
		panic(err)
	}
	defer os.Remove(path)

	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		panic(err)
	}
	err = syscall.SetFileAttributes(p, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		panic(err)
	}
}
