package ffh

import "os"

type GoFile struct {
	Path    string
	Content string
}

func NewGoFile(path string) (*GoFile, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	goFile := &GoFile{
		Path:    path,
		Content: string(fileContent),
	}
	return goFile, nil
}
