package util

import (
	"errors"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var sourceDir string

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get source directory with various operating systems
	sourceDir = regexp.MustCompile(`util.util\.go`).ReplaceAllString(file, "")
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	shouldStop := func(file string) bool {
		if !strings.HasPrefix(file, sourceDir) {
			return false
		}

		file = file[len(sourceDir):]

		excludedPrefix := []string{"util", "internal", "repository/repository.go"}
		for _, prefix := range excludedPrefix {
			if strings.HasPrefix(file, prefix) {
				return false
			}
		}

		return true
	}

	// the first caller usually from here, so set i start from 1
	for i := 1; i < 20; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && shouldStop(file) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}
