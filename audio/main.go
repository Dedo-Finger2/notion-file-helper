package audio

import (
	"io"
	"os"
	"strings"

	"github.com/Dedo-Finger2/notion-file-helper/utils"
)

func CutMp3AudioInHalf(path string) error {
	const (
		FILE_FORMAT        = ".mp3"
		FILE_PERMISSIONS   = 0644 // 0 (octal), 6 (read and write), 4 (group read), 4 (other read)
		PART_ONE_FILE_NAME = "part1" + FILE_FORMAT
		PART_TWO_FILE_NAME = "part2" + FILE_FORMAT
	)

	zipName := strings.Split(strings.Split(path, "/")[len(strings.Split(path, "/"))-1], ".")[0]

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	s := info.Size()
	hs := s / 2

	buf := make([]byte, s)

	_, err = io.ReadFull(f, buf)
	if err != nil {
		return err
	}

	os.WriteFile(PART_ONE_FILE_NAME, buf[:hs], FILE_PERMISSIONS)
	os.WriteFile(PART_TWO_FILE_NAME, buf[hs:], FILE_PERMISSIONS)

	utils.CreateZipFileWithFiles(zipName, []string{PART_ONE_FILE_NAME, PART_TWO_FILE_NAME})

	return nil
}
