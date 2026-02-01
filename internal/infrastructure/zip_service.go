package video

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func CreateZip(files []string, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	for _, file := range files {
		f, _ := os.Open(file)
		defer f.Close()

		w, _ := writer.Create(filepath.Base(file))
		io.Copy(w, f)
	}
	return nil
}
