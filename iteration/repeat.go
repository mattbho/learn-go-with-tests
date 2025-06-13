package iteration

import "strings"

// Repeats a character repeatCount times
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func StringRepeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
