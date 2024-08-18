package ffh

import (
	"io/fs"
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

// creates a file or dir
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
func CollectFilesDownward(dir string) ([]string, error) {
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

// return a go import statement
func GoImport(imports []string) (string, error) {
	if len(imports) == 0 {
		return "", nil
	}
	if len(imports) == 1 {
		return "import \"" + imports[0] + "\"\n", nil
	}
	content := "import (\n"
	for _, imp := range imports {
		content += "\t\"" + imp + "\"\n"
	}
	content += ")\n"
	return content, nil
}
