package ffh

import "strings"

// work on lines of a file, line by line, getting a new string for each line
func MapFileLines(path string, f func(string) string) ([]string, error) {
	content, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(content, "\n")
	newLines := make([]string, len(lines))
	for i, line := range lines {
		newLines[i] = f(line)
	}
	return newLines, nil
}

// work on each string in a slice, returning a new slice
func MapStrings(ss []string, f func(int, string) string) []string {
	newStrings := make([]string, 0)
	for i, s := range ss {
		newString := f(i, s)
		if newString == "" {
			continue
		}
		newStrings = append(newStrings, newString)
	}
	return newStrings
}
