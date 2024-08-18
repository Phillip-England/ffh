package ffh

import "strings"

func ExtractGoFunc(input string, funcName string) (string, error) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.Contains(line, "func "+funcName) {
			for j := i; j < len(lines); j++ {
				if lines[j] == "}" {
					return strings.Join(lines[i:j+1], "\n"), nil
				}
			}
		}
	}
	return "", nil
}

func ExtractGoTypeFunc(input string, typeName string, funcName string) (string, error) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.Contains(line, "func") && strings.Contains(line, typeName) && strings.Contains(line, funcName) {
			for j := i; j < len(lines); j++ {
				if lines[j] == "}" {
					return strings.Join(lines[i:j+1], "\n"), nil
				}
			}
		}
	}
	return "", nil
}

func ExtractGoType(input string, typeName string) (string, error) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.Contains(line, "type "+typeName) {
			for j := i; j < len(lines); j++ {
				if lines[j] == "}" {
					return strings.Join(lines[i:j+1], "\n"), nil
				}
			}
		}
	}
	return "", nil
}
