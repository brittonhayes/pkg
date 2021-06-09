// Package check provides simple utilities
// to check the status or state of things
package check

import (
	"go/format"
	"os"
)

// FileExists searches for the provided filename
// and checks if it is present.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FilePerms searches for the provided filename
// and checks if it is present and returns the file's
// permissions.
func FilePerms(filename string) (*os.FileMode, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, err
	}
	perm := info.Mode().Perm()
	return &perm, nil
}

// FolderExists searches for the provided directory
// name and checks if it is present.
func FolderExists(folder string) bool {
	info, err := os.Stat(folder)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// IsGoFormatted checks if a Go file's contents
// are formatted using go fmt
func IsGoFormatted(b []byte) bool {
	dst, err := format.Source(b)
	if err != nil {
		return false
	}
	if string(b) != string(dst) {
		return false
	}
	return true
}
