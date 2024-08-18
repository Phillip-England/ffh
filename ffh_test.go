package ffh

import (
	"fmt"
	"strings"
	"testing"
)

//======================================
// testing i/o operations
//======================================

const FILE = "ffh_doc.txt"

func TestBasicIO(t *testing.T) {
	err := WriteFile(FILE, "Hello, World!")
	if err != nil {
		t.Errorf("WriteFile() failed: %v", err)
	}
	content, err := ReadFile(FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
	err = ClearFile(FILE)
	if err != nil {
		t.Errorf("ClearFile() failed: %v", err)
	}
	content, err = ReadFile(FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
	if !FileExists(FILE) {
		t.Errorf("!FileExists() failed: file should exist")
	}
	err = WriteFileLines(FILE, []string{"Hello", "World"})
	if err != nil {
		t.Errorf("WriteFileLines() failed: %v", err)
	}
	content, err = ReadFile(FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello\nWorld" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
	err = AppendFile(FILE, "!")
	if err != nil {
		t.Errorf("AppendFile() failed: %v", err)
	}
	content, err = ReadFile(FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello\nWorld!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
	err = AppendFileLine(FILE, "!")
	if err != nil {
		t.Errorf("AppendFileLine() failed: %v", err)
	}
	content, err = ReadFile(FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello\nWorld!\n!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

//======================================
// testing higher order functions
//======================================

func TestMapFileLines(t *testing.T) {
	err := WriteFile(FILE, "Hello, World!")
	if err != nil {
		t.Errorf("WriteFile() failed: %v", err)
	}
	lines, err := MapFileLines(FILE, strings.ToUpper)
	if err != nil {
		t.Errorf("MapFileLines() failed: %v", err)
	}
	if lines[0] != "HELLO, WORLD!" {
		t.Errorf("MapFileLines() failed: content mismatch")
	}
	err = WriteFileLines(FILE, lines)
	if err != nil {
		t.Errorf("WriteFileLines() failed: %v", err)
	}
	content, err := ReadFile(FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "HELLO, WORLD!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

//======================================
// testing writing go code
//======================================

func TestGoType(t *testing.T) {
	goType, err := GoType(FILE, "Person", []string{
		"Name string",
		"Age int",
	})
	if err != nil {
		t.Errorf("GoType() failed: %v", err)
	}
	fmt.Println(goType)
}

func TestGoFunc(t *testing.T) {
	goFunc, err := GoFunc(FILE, "Hello", "name string", "string", `
		return "Hello, " + name + "!"
	`)
	if err != nil {
		t.Errorf("GoFunc() failed: %v", err)
	}
	fmt.Println(goFunc)
}

func TestGoTypeFunc(t *testing.T) {
	goTypeFunc, err := GoTypeFunc(FILE, "p *Person", "Greet", "name string", "string", `
		return "Hello, " + p.Name + "!"
	`)
	if err != nil {
		t.Errorf("GoTypeFunc() failed: %v", err)
	}
	fmt.Println(goTypeFunc)
}
