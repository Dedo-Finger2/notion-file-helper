package utils

import (
	"archive/zip"
	"io"
	"os"
)

// Returns size in bytes, megabytes and a possible error.
func GetFileSizeFromPath(path string) (int64, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil {
		return 0, 0, err
	}
	return fInfo.Size(), float64(fInfo.Size()) / float64(1000000), nil
}

func CreateZipFileWithFiles(name string, files []string) error {
	zipFile, err := os.Create(name + ".zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for _, file := range files {
		srcFile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		zipEntry, err := zipWriter.Create(file)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipEntry, srcFile)
		if err != nil {
			return err
		}
		os.Remove(file)
	}
	return nil
}
