package util

import (
	"runtime"
	"strings"
)

const (
	WindowsNewLine string = "\r\n"
	UnixNewLine    string = "\n"
)

// Function for removing all new line characters from a string input.
func RemoveAllNewLineChars(input string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.ReplaceAll(input, WindowsNewLine, "")
	default:
		return strings.ReplaceAll(input, UnixNewLine, "")
	}
}
