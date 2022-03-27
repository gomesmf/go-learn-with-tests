package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}
