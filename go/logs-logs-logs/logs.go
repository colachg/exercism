package logs

import (
	"fmt"
	"strings"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	return strings.TrimSpace(strings.Split(line, ":")[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	var chars []rune
	for _, c := range Message(line) {
		chars = append(chars, c)
	}
	return len(chars)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	s := strings.Split(line, ":")[0]
	s = strings.Trim(s, "[]")
	return strings.ToLower(s)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
