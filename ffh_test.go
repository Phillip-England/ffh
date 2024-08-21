package ffh

import (
	"strings"
	"testing"
)

// a type for testing only
type FuncDef string

// another testing type
type Route struct {
	Method string
}

const PATH = "./ffh_test.go"

func Test_ExtractPackageLine(t *testing.T) {
	str, _ := ReadFile(PATH)
	packageLine, _ := ExtractPackageLine(str)
	if packageLine != "package ffh" {
		t.Error("failed to extract package correctly")
	}
}

func Test_ExtractTypeBlocks(t *testing.T) {
	str, _ := ReadFile(PATH)
	goTypes, _ := ExtractTypeBlocks(str)
	if len(goTypes) != 2 {
		t.Error("expected 2 test types to be found in this file")
	}
}

func Test_ExtractFuncBlocks(t *testing.T) {
	str, _ := ReadFile(PATH)
	goFuncs, _ := ExtractFuncBlocks(str)
	if len(goFuncs) != 4 {
		t.Error("expected 4 func blocks in this file")
	}
}

func Test_ExtractImportBlock(t *testing.T) {
	str, _ := ReadFile(PATH)
	block, _ := ExtractImportBlock(str)
	if !strings.Contains(block, "testing") {
		t.Error("failed to extract import block correctly")
	}
}
