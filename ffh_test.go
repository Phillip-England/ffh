package ffh

import (
	"fmt"
	"testing"
)

type Route struct {
	Method string
}

func (r *Route) Render() {
	fmt.Println(r.Method)
}

func TestGoFile(t *testing.T) {
	f, _ := NewGoFile("ffh_test.go")
	for _, imp := range f.GoTypes {
		fmt.Println(imp.Name)
		// fmt.Println(imp.Content)
	}
}
