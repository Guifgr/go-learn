package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func main() {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	result, _, _ := transform.String(t, "\U0001F601")

}
