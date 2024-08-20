package ffh

import (
	"fmt"
	"os"
	"strings"
)

type GoFile struct {
	PackageName string
	Path        string
	GoFuncs     []*GoFunc
	GoImport    *GoImport
}

func NewGoFile(path string) (*GoFile, error) {
	f := &GoFile{
		Path: path,
	}
	err := f.Init()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *GoFile) Init() error {
	funcs, err := f.ExtractFuncs()
	if err != nil {
		return err
	}
	for _, fn := range funcs {
		goFunc, err := NewGoFunc(fn)
		if err != nil {
			return nil
		}
		f.GoFuncs = append(f.GoFuncs, goFunc)
	}
	packageStr, err := f.ExtractPackage()
	if err != nil {
		return err
	}
	packageSplit := strings.Split(packageStr, " ")
	if len(packageSplit) != 2 {
		return fmt.Errorf("provided go file is missing a valid package name")
	}
	packageName := packageSplit[1]
	f.PackageName = packageName
	importStr, err := f.ExtractImport()
	if err != nil {
		return nil
	}
	goImport, err := NewGoImport(importStr)
	if err != nil {
		return err
	}
	f.GoImport = goImport
	return nil
}

func (f *GoFile) Read() (string, error) {
	t, err := os.ReadFile(f.Path)
	if err != nil {
		return "", err
	}
	return string(t), nil
}

func (f *GoFile) ExtractImport() (string, error) {
	t, err := f.Read()
	if err != nil {
		return "", err
	}
	importLines := []string{}
	foundImport := false
	for _, line := range strings.Split(t, "\n") {
		chunckLen := 6
		chunck := line
		if len(line) > chunckLen {
			chunck = line[:chunckLen]
		}
		if chunck == "import" {
			foundImport = true
			chunck2 := line[:chunckLen+2]
			if chunck2 == "import \"" {
				return line, nil
			}
		}
		if foundImport {
			importLines = append(importLines, line)
			if line == ")" {
				break
			}
		}
	}
	return strings.Join(importLines, "\n"), nil
}

func (f *GoFile) ExtractPackage() (string, error) {
	t, err := f.Read()
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(t, "\n") {
		if strings.Contains(line, "package") {
			return line, nil
		}
	}
	return "", fmt.Errorf("go file does not contain package")
}

func (f *GoFile) ExtractFuncs() ([]string, error) {
	blocks := []string{}
	t, err := f.Read()
	if err != nil {
		return blocks, err
	}
	lines := strings.Split(t, "\n")
	currentFunc := ""
	inFuncBlock := false
	for _, line := range lines {
		if strings.Contains(line, "func") && strings.Contains(line, "(") && strings.Contains(line, ") {") {
			inFuncBlock = true
		}
		if !inFuncBlock {
			continue
		}
		currentFunc = currentFunc + line + "\n"
		if line == "}" {
			inFuncBlock = false
			blocks = append(blocks, currentFunc)
			currentFunc = ""
		}
	}
	finalBlocks := []string{}
	// skipping funcs on types (could easily collect them here as well)
	for _, block := range blocks {
		firstLine := strings.Split(block, "\n")[0]
		alteredLine := strings.Replace(firstLine, "func ", "", 1)
		if alteredLine[0] == '(' {
			continue
		}
		finalBlocks = append(finalBlocks, block)
	}
	return finalBlocks, nil
}

func (f *GoFile) ExtractTypes() ([]string, error) {
	blocks := []string{}
	t, err := f.Read()
	if err != nil {
		return blocks, err
	}
	lines := strings.Split(t, "\n")
	currentBlock := ""
	inTypeBlock := false
	for _, line := range lines {
		if strings.Contains(line, "type") && strings.Contains(line, "{") {
			if strings.Contains(line, "{") {
				inTypeBlock = true
				continue
			}
			blocks = append(blocks, line)
			return blocks, nil
		}
		if !inTypeBlock {
			continue
		}
		currentBlock = currentBlock + line + "\n"
		if line == "}" {
			inTypeBlock = false
			blocks = append(blocks, currentBlock)
			currentBlock = ""
		}
	}
	return blocks, nil
}
