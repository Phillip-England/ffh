package ffh

import (
	"strings"
	"testing"
)

//======================================
// testing i/o operations
//======================================

const PATH = "ffh_doc.txt"

func TestClearFile(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = ClearFile(PATH)
	if err != nil {
		t.Errorf("ClearFile() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestAppendFile(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = AppendFile(PATH, "!")
	if err != nil {
		t.Errorf("AppendFile() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestAppendFileLine(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = AppendFileLine(PATH, "!")
	if err != nil {
		t.Errorf("AppendFileLine() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World\n!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestFileExists(t *testing.T) {
	if !FileExists(PATH) {
		t.Errorf("!FileExists() failed: file should exist")
	}
}

func TestOverwriteFile(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestReadFile(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestOverwriteFileLines(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = OverwriteFileLines(PATH, []string{"Hello", "World"})
	if err != nil {
		t.Errorf("OverwriteFileLines() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello\nWorld" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestCollectFilesDownward(t *testing.T) {
	files, err := CollectFilesDownward(".")
	if err != nil {
		t.Errorf("CollectFilesDownward() failed: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("CollectFilesDownward() failed: no files found")
	}
	goFiles := MapStrings(files, func(_ int, s string) string {
		if strings.HasSuffix(s, ".go") {
			return s
		}
		return ""
	})
	if !SliceContains(goFiles, "ffh_test.go") {
		t.Errorf("CollectFilesDownward() failed: go files not found")
	}
}

func TestSliceContains(t *testing.T) {
	ss := []string{"Hello", "World"}
	if !SliceContains(ss, "Hello") {
		t.Errorf("SliceContains() failed: should contain 'Hello'")
	}
}

func TestRemoveFile(t *testing.T) {
	err := RemoveFile("./test")
	if err != nil {
		t.Errorf("RemoveFile() failed: %v", err)
	}
	if FileExists("./test") {
		t.Errorf("RemoveFile() failed: file not removed")
	}
}

func TestTouch(t *testing.T) {
	err := Touch("test")
	if err != nil {
		t.Errorf("Mkdir() failed: %v", err)
	}
	if !FileExists("test") {
		t.Errorf("Mkdir() failed: directory not created")
	}
}

//======================================
// testing higher order functions
//======================================

func TestMapFileLines(t *testing.T) {
	err := OverwriteFile(PATH, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	lines, err := MapFileLines(PATH, strings.ToUpper)
	if err != nil {
		t.Errorf("MapFileLines() failed: %v", err)
	}
	if lines[0] != "HELLO, WORLD!" {
		t.Errorf("MapFileLines() failed: content mismatch")
	}
	err = OverwriteFileLines(PATH, lines)
	if err != nil {
		t.Errorf("OverwriteFileLines() failed: %v", err)
	}
	content, err := ReadFile(PATH)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "HELLO, WORLD!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestMapStrings(t *testing.T) {
	lines := []string{"Hello", "World"}
	newLines := MapStrings(lines, func(_ int, s string) string {
		return strings.ToUpper(s)
	})
	if newLines[0] != "HELLO" {
		t.Errorf("MapStrings() failed: content mismatch")
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
