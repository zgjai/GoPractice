package main

import (
	"unicode"
	"strings"
	"fmt"
)

func main() {
	s := "abcd ef"
	fmt.Println(s, title(s))
}

func title(s string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				fmt.Println("true", string(prev), string(r), &prev, &r)
				prev = r
				fmt.Println("true", string(prev), string(r), &prev, &r)
				return unicode.ToTitle(r)
			}
			fmt.Println(string(prev), string(r), &prev, &r)
			prev = r
			fmt.Println(string(prev), string(r), &prev, &r)
			return r
		},
		s)
}

func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}