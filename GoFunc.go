package ffh

import (
	"strings"
)

type GoFunc struct {
	Content string
	Body    string
	Name    string
}

func NewGoFunc(goFuncStr string) (*GoFunc, error) {
	goFunc := &GoFunc{
		Content: goFuncStr,
	}
	err := goFunc.Init()
	if err != nil {
		return nil, err
	}
	return goFunc, nil
}

func (fn *GoFunc) Init() error {
	body := fn.ExtractBody()
	fn.Body = body
	name := fn.ExtractName()
	fn.Name = name
	return nil
}

func (fn *GoFunc) ExtractBody() string {
	if len(strings.Split(fn.Content, "\n")) <= 2 {
		return ""
	}
	bodyStr := ""
	for i, line := range strings.Split(fn.Content, "\n") {
		if i == 0 {
			continue
		}
		if i == len(strings.Split(fn.Content, "\n"))-2 {
			break
		}
		bodyStr = bodyStr + line + "\n"
	}
	return bodyStr
}

func (fn *GoFunc) ExtractName() string {
	firstLine := strings.Split(fn.Content, "\n")[0]
	firstLine = strings.Replace(firstLine, "func ", "", 1)
	for i, char := range firstLine {
		if char == '(' {
			firstLine = firstLine[:i]
		}
	}
	return firstLine
}
