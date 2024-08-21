package ffh

import (
	"fmt"
	"os"
	"strings"
)

// mocks a range loop; if you return true in the loop, it will break
func LoopLines(str string, fn func(i int, line string) bool) {
	lines := strings.Split(str, "\n")
	for i, ln := range lines {
		shouldBreak := fn(i, ln)
		if shouldBreak {
			break
		}
	}
}

// general loop with generics, if you return true, it will break
func Loop[T any](items []T, fn func(i int, item T) bool) {
	for i, itm := range items {
		shouldBreak := fn(i, itm)
		if shouldBreak {
			break
		}
	}
}

// read a file to a string
func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// extracts all the funcs out of a .go file string
func ExtractFuncBlocks(str string) ([]string, error) {
	funcs := []string{}
	err := StrIsGoFile(str)
	if err != nil {
		return funcs, err
	}
	currentFunc := ""
	LoopLines(str, func(i int, line string) bool {
		if strings.Contains(line, "func ") && strings.Contains(line, "(") && strings.Contains(line, ")") && strings.Contains(line, "{") {
			LoopLines(str, func(i2 int, line2 string) bool {
				if i2 >= i {
					currentFunc = currentFunc + line2 + "\n"
					if line2 == "}" {
						funcs = append(funcs, currentFunc)
						currentFunc = ""
						return true
					}
				}
				return false
			})
		}
		return false
	})
	return funcs, nil
}

// extracts the import statement out of a .go file string
func ExtractImportBlock(str string) (string, error) {
	err := StrIsGoFile(str)
	if err != nil {
		return "", err
	}
	statement := ""
	LoopLines(str, func(i int, line string) bool {
		if strings.Contains(line, "import \"") {
			statement = line
			return true
		}
		if strings.Contains(line, "import (") {
			LoopLines(str, func(i2 int, line2 string) bool {
				if i2 >= i {
					statement = statement + line2 + "\n"
					if line2 == ")" {
						return true
					}
				}
				return false
			})
		}
		return false
	})
	return statement, nil
}

// extracts all the types out of a .go file string
func ExtractTypeBlocks(str string) ([]string, error) {
	typeBlocks := []string{}
	err := StrIsGoFile(str)
	if err != nil {
		return typeBlocks, err
	}
	currentType := ""
	LoopLines(str, func(i int, line string) bool {
		if strings.Contains(line, "type ") {
			if !strings.Contains(line, "{") {
				typeBlocks = append(typeBlocks, line)
				return false
			}
			LoopLines(str, func(i2 int, line2 string) bool {
				if i2 >= i {
					currentType = currentType + line2 + "\n"
					if line2 == "}" {
						typeBlocks = append(typeBlocks, currentType)
						currentType = ""
						return true
					}
				}
				return false
			})
		}
		return false
	})
	return typeBlocks, nil
}

// extracts the package line out of a go file string
func ExtractPackageLine(str string) (string, error) {
	err := StrIsGoFile(str)
	if err != nil {
		return "", err
	}
	packageLine := ""
	LoopLines(str, func(i int, line string) bool {
		if strings.Contains(line, "package ") {
			packageLine = line
			return true
		}
		return false
	})
	if packageLine == "" {
		return packageLine, fmt.Errorf("package line not found in .go file string")
	}
	return packageLine, nil
}

// checks if a provided string is a .go file
func StrIsGoFile(str string) error {
	foundTypeDef := false
	foundFuncDef := false
	foundImport := false
	foundPackage := false
	if strings.Contains(str, "package ") {
		foundPackage = true
	}
	if strings.Contains(str, "func ") {
		foundFuncDef = true
	}
	if strings.Contains(str, "type ") {
		foundTypeDef = true
	}
	if strings.Contains(str, "import ") {
		foundImport = true
	}
	if !foundPackage {
		return fmt.Errorf("provided str did not contain a package definition")
	}
	if !foundImport {
		return fmt.Errorf("provided str did not contain an import statement")
	}
	if !foundFuncDef && !foundTypeDef {
		return fmt.Errorf("provided str did not contain any type definitions or func definitions")
	}
	return nil
}
