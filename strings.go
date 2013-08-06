package kagus

import (
	"strings"
	"unicode"
)

// StringConsistsOf tests, if f returns true for all runes in s
func StringConsistsOf(s string, f func(rune) bool) bool {
	for _, r := range s {
		if !f(r) {
			return false
		}
	}
	return true
}

// IsASCII tests, if a Rune is in the ASCII set
func IsASCII(r rune) bool {
	return (r <= unicode.MaxASCII)
}

// Indent a string s by prefixing each line with ind.
func Indent(s string, ind string) (rv string) {
	ls := strings.Split(s, "\n")
	first := true
	for _, l := range ls {
		if !first {
			rv += "\n"
		}
		rv += ind + l
		first = false
	}
	return
}
