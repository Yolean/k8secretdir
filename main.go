package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <direcotry path>\n", os.Args[0])
		os.Exit(2)
	}
	path := os.Args[1]
	fmt.Printf(`---
apiVersion: v1
kind: Secret
metadata:
  name: %s
type: Opaque
data:
`, filepath.Base(path))
	if err := filepath.Walk(path, encode); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func encode(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("  %s: >\n    ", filepath.Base(path))
	enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer fmt.Println()
	defer enc.Close()
	_, err = io.Copy(enc, f)
	if err != nil {
		return err
	}
	return nil
}
