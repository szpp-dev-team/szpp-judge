package fsutil

import (
	"os"
	"path"
	"testing"
)

func TestGetGoModAbsDir(t *testing.T) {
	dir := GetGoModAbsDir()
	file := path.Join(dir, "go.mod")

	_, err := os.Stat(file)
	if err != nil {
		t.Fatalf("Not a go.mod directory: %s", dir)
	}

	if !path.IsAbs(dir) {
		t.Fatalf("Expected abs path, but got %s", dir)
	}
}
