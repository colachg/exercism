// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isQuestion(text string) bool {
	text = strings.TrimSpace(text)
	if strings.HasSuffix(text,"?") {
		return  true
	}
	return  false
}

func isShouting(text string) bool {
	text = strings.TrimSpace(text)
	if strings.ToUpper(text) == text &&  strings.IndexFunc(text, unicode.IsLetter) >= 0 {
		return true
	}
	return false

}
func isEmpty(text string) bool {
	text = strings.TrimSpace(text)
	if text == "" {
		return  true
	}
	return false
}


// Hey should have a comment documenting it.
func Hey(remark string) string {
	switch {
	// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
	case isShouting(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
	case isShouting(remark):
		return "Whoa, chill out!"
	// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
	case isQuestion(remark):
		return "Sure."
	// He says 'Fine. Be that way!' if you address him without actually saying anything.
	case isEmpty(remark):
		return "Fine. Be that way!"
	// He answers 'Whatever.' to anything else
	default:
		return "Whatever."
	}
}
