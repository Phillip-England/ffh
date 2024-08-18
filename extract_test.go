package ffh

import (
	"fmt"
	"testing"
)

type TestType struct {
	Name string
	Age  int
}

func (t TestType) FuncTest() {
	fmt.Println("Hello, World!")
}

func TestExtractGoFunc(t *testing.T) {
	path := "./extract_test.go"
	txt, err := ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	_, err = ExtractGoFunc(txt, "TestExtractGoFunc")
	if err != nil {
		t.Errorf("ExtractGoFunc() failed: %v", err)
	}
	// fmt.Println(goFunc)
}

func TestExtractGoTypeFunc(t *testing.T) {
	path := "./extract_test.go"
	txt, err := ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	_, err = ExtractGoTypeFunc(txt, "TestType", "FuncTest")
	if err != nil {
		t.Errorf("ExtractGoTypeFunc() failed: %v", err)
	}
	// fmt.Println(goFunc)
}

func TestExtractGoType(t *testing.T) {
	path := "./extract_test.go"
	txt, err := ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	_, err = ExtractGoType(txt, "TestType")
	if err != nil {
		t.Errorf("ExtractGoType() failed: %v", err)
	}
}
