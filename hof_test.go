package ffh

import (
	"strings"
	"testing"
)

func TestMapFileLines(t *testing.T) {
	err := OverwriteFile(TEST_FILE, "Hello, World!")
	if err != nil {
		t.Errorf("OverwriteFile() failed: %v", err)
	}
	lines, err := MapFileLines(TEST_FILE, strings.ToUpper)
	if err != nil {
		t.Errorf("MapFileLines() failed: %v", err)
	}
	if lines[0] != "HELLO, WORLD!" {
		t.Errorf("MapFileLines() failed: content mismatch")
	}
	err = OverwriteFileLines(TEST_FILE, lines)
	if err != nil {
		t.Errorf("OverwriteFileLines() failed: %v", err)
	}
	content, err := ReadFile(TEST_FILE)
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
