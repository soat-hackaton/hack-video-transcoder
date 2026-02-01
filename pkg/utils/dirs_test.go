package utils

import (
	"os"
	"testing"
)

func TestCreateDirs(t *testing.T) {
	dir := "uploads"

	err := CreateRequiredDirs()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("directory was not created")
	}
}
