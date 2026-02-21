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

		info, _ := f.Stat()
		header, _ := zip.FileInfoHeader(info)
		header.Name = filepath.Base(file)
		header.Method = zip.Store

		w, _ := writer.CreateHeader(header)
		io.Copy(w, f)
	}
	return nil
}
