package ffh

import (
	"strings"
	"testing"
)

func TestTouch(t *testing.T) {
	err := Touch("test")
	if err != nil {
		t.Errorf("Mkdir() failed: %v", err)
	}
	if !FileExists("test") {
		t.Errorf("Mkdir() failed: directory not created")
	}
}

func TestClearFile(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = ClearFile(TEST_FILE)
	if err != nil {
		t.Errorf("ClearFile() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestAppendFile(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = AppendFile(TEST_FILE, "!")
	if err != nil {
		t.Errorf("AppendFile() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestAppendFileLine(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = AppendFileLine(TEST_FILE, "!")
	if err != nil {
		t.Errorf("AppendFileLine() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World\n!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestFileExists(t *testing.T) {
	if !FileExists(TEST_FILE) {
		t.Errorf("!FileExists() failed: file should exist")
	}
}

func TestOverwriteFile(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestReadFile(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello, World!" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestOverwriteFileLines(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	err = OverwriteFileLines(TEST_FILE, []string{"Hello", "World"})
	if err != nil {
		t.Errorf("OverwriteFileLines() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if content != "Hello\nWorld" {
		t.Errorf("ReadFile() failed: content mismatch")
	}
}

func TestCollectFilesCascade(t *testing.T) {
	files, err := CollectFilesCascade(".")
	if err != nil {
		t.Errorf("CollectFilesCascade() failed: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("CollectFilesCascade() failed: no files found")
	}
	goFiles := MapStrings(files, func(_ int, s string) string {
		if strings.HasSuffix(s, ".go") {
			return s
		}
		return ""
	})
	if !SliceContains(goFiles, "const.go") {
		t.Errorf("CollectFilesCascade() failed: go files not found")
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
