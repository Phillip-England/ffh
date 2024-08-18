package ffh

import (
	"os"
	"strings"
)

//======================================
// i/o operations
//======================================

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

//======================================
// higher order functions
//======================================

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

//======================================
// writing go code
//======================================

// return a go type definition
func GoType(typeName string, fields []string) (string, error) {
	content := "type " + typeName + " struct {\n"
	for _, field := range fields {
		content += "\t" + field + "\n"
	}
	content += "}\n"
	return content, nil
}

// return a go function definition
func GoFunc(name string, params string, returnStr string, body string) (string, error) {
	content := "func " + name + "(" + params + ") " + returnStr + " {\n"
	content += body
	content += "}\n"
	contentLines := strings.Split(content, "\n")
	newLines := make([]string, 0)
	for i, line := range contentLines {
		newLine := strings.Replace(line, "\t", "", 1)
		if i == 1 {
			continue
		}
		newLines = append(newLines, newLine)
	}
	content = strings.Join(newLines, "\n")
	return content, nil
}

// return a go func for a type
func GoTypeFunc(typeName string, name string, params string, returnStr string, body string) (string, error) {
	content := "func (" + typeName + ") " + name + "(" + params + ") " + returnStr + " {\n"
	content += body
	content += "}\n"
	contentLines := strings.Split(content, "\n")
	newLines := make([]string, 0)
	for i, line := range contentLines {
		newLine := strings.Replace(line, "\t", "", 1)
		if i == 1 {
			continue
		}
		newLines = append(newLines, newLine)
	}
	content = strings.Join(newLines, "\n")
	return content, nil
}
