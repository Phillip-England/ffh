package ffh

import (
	"fmt"
	"strings"
)

type GoImport struct {
	Content string
	Imports []string
}

func NewGoImport(importStr string) (*GoImport, error) {
	goImport := &GoImport{
		Content: importStr,
	}
	err := goImport.Init()
	if err != nil {
		return nil, err
	}
	return goImport, nil
}

func (i *GoImport) Init() error {
	imports, err := i.ExtractImports()
	if err != nil {
		return err
	}
	i.Imports = imports
	return nil
}

func (i *GoImport) ExtractImports() ([]string, error) {
	lines := strings.Split(i.Content, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("invalid import statement provided")
	}
	imports := []string{}
	if strings.Contains(lines[0], "\"") {
		imports = append(imports, strings.Split(lines[0], " ")[1])
		return imports, nil
	}
	for _, line := range lines {
		if strings.Contains(line, "(") || strings.Contains(line, ")") {
			continue
		}
		imports = append(imports, line)
	}
	return imports, nil
}
