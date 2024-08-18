package ffh

import (
	"io/fs"
	"os"
	"strings"
)

// ClearFile clears the file
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// append content to a file
func AppendFile(path string, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// append a line to a file
func AppendFileLine(path string, line string) error {
	return AppendFile(path, "\n"+line)
}

// appends lines to a file
func AppendFileLines(path string, lines []string) error {
	content := strings.Join(lines, "\n")
	return AppendFile(path, content)
}

// writes content to a file - overwrites existing content - creates file if it does not exist
func OverwriteFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}
	return nil
}

// reads content from a file
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// writes lines to a file
func OverwriteFileLines(path string, lines []string) error {
	content := strings.Join(lines, "\n")
	return OverwriteFile(path, content)
}

// creates a file or dir if it does not exist
func Touch(dir string) error {
	err := os.Mkdir(dir, 0777)
	if err != nil {
		return err
	}
	return nil
}

// removes a file or dir
func RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// take in a dir name, walk through it and its children, and collect the paths of all files
func CollectFilesCascade(dir string) ([]string, error) {
	var paths []string
	err := fs.WalkDir(os.DirFS(dir), dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// checks if a slice contains a string
func SliceContains(ss []string, s string) bool {
	for _, str := range ss {
		if str == s {
			return true
		}
	}
	return false
}

// takes a slice of lines and returns a string
func LinesToStr(lines []string) string {
	return strings.Join(lines, "\n")
}
