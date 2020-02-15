package iteration

import "strings"

const defaultRepeatCount = 5

// Repeat takes a character and returns it repeated.
func Repeat(character string, repeatCount int) string {
	if repeatCount <= 0 {
		repeatCount = defaultRepeatCount
	}

	return strings.Repeat(character, repeatCount)
}
