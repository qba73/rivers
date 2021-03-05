package testhelper

import (
	"io/ioutil"
	"os"
	"testing"
)

// TmpFile knows hwo to create temporary file for testing.
func TmpFile(t *testing.T, dirname, filename string) *os.File {
	file, err := ioutil.TempFile(dirname, filename)
	if err != nil {
		t.Fatalf("failure to create a temporary test file: %s", err)
	}

	return file
}

// TmpTextFile knows how to create temporary text file
// with provided content.
func TmpTextFile(t *testing.T, dirname, filename, content string) *os.File {
	file := TmpFile(t, dirname, filename)
	_, err := file.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write to temp file: %s", err)
	}

	return file
}
