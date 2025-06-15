package utils

import "os"

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
