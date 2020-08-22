package replayd

import "os"

func OpenWritableFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_CREATE, 0555)
	if err != nil {
		return nil, err
	}
	return file, nil
}
