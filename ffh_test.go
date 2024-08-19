package ffh

import (
	"fmt"
	"testing"
)

func TestFFH(t *testing.T) {
	goFile, _ := NewGoFile("ffh_test.go")
	fmt.Println(goFile.Content)
}

// newfunc
