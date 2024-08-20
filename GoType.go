package ffh

import (
	"fmt"
	"strings"
)

type GoType struct {
	Content string
	Name    string
}

func NewGoType(goTypeStr string) (*GoType, error) {
	goType := &GoType{
		Content: goTypeStr,
	}
	// err := goType.Init()
	// if err != nil {
	// 	return nil, err
	// }
	return goType, nil
}

func (t *GoType) Init() error {
	lines := strings.Split(t.Content, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("invalid go type provided")
	}
	firstLine := lines[0]
	parts := strings.Split(firstLine, " ")
	if len(parts) < 3 {
		return fmt.Errorf("invalid go type provided")
	}
	// t.Name = parts[1]
	return nil
}
