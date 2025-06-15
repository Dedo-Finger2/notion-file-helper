package audio

import (
	"os"
)

func GetAudioFileSizeFromPath(path string) (float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return float64(fInfo.Size()) / float64(1000000), nil
}
