package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Copy(source, dest string) error {
	sfs, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !sfs.Mode().IsRegular() {
		return fmt.Errorf("source file: '%s' is not a regular file.", source)
	}

	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dst.Close()
	nbts, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	fmt.Printf("file copied to '%s' size: %d bytes.\n", dest, nbts)
	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	// Also check if it's a directory (optional, depending on requirements)
	return err == nil && !info.IsDir()
}
