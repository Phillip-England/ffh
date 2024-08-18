package ffh

import (
	"strings"
	"testing"
)

//======================================
// testing i/o operations
//======================================

const FILE = "ffh_doc.txt"

func TestBasicIO(t *testing.T) {
	err := OverwriteFile(FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
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
	err = OverwriteFileLines(FILE, []string{"Hello", "World"})
	if err != nil {
		t.Errorf("OverwriteFileLines() failed: %v", err)
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
	err := OverwriteFile(FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	lines, err := MapFileLines(FILE, strings.ToUpper)
	if err != nil {
		t.Errorf("MapFileLines() failed: %v", err)
	}
	if lines[0] != "HELLO, WORLD!" {
		t.Errorf("MapFileLines() failed: content mismatch")
	}
	err = OverwriteFileLines(FILE, lines)
	if err != nil {
		t.Errorf("OverwriteFileLines() failed: %v", err)
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

}

func TestGoFunc(t *testing.T) {

}

func TestGoTypeFunc(t *testing.T) {

}
