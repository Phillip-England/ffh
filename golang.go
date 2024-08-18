package ffh

import "strings"

// build a go struct
func GoStruct(typeName string, fields []string) (string, error) {
	content := "type " + typeName + " struct {\n"
	for _, field := range fields {
		content += "\t" + field + "\n"
	}
	content += "}\n"
	return content, nil
}

// build a go function definition
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

// build a go type function definition
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

// build a go import statement
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
