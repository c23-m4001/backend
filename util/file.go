package util

import (
	"errors"
	"os"
)

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
