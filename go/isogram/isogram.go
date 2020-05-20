// Package isogram detects isograms
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if the input s is an isogram
func IsIsogram(s string) bool {
	s = strings.ToLower(s)

	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}

	return true
}
