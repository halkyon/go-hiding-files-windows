package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := os.TempDir()
	path := filepath.Join(dir, "testfile.txt")
	err := ioutil.WriteFile(path, []byte("hello\n"), 0600)
	if err != nil {
		panic(err)
	}

	p, err := HideFile(path)
	if err != nil {
		panic(err)
	}

	os.Remove(p)
	fmt.Println(p)
}
