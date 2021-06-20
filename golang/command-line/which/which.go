package which

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrExecutableNotFound = errors.New("executable not found")

func Which(file string) (string, error) {
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		// Is it a regular file and executable?
		mode := fileInfo.Mode()
		if mode.IsRegular() && mode&0111 != 0 {
			return fullPath, nil
		}
	}

	return "", ErrExecutableNotFound
}
