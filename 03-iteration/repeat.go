package iteration

import "strings"

func Repeat(character string, count int) string {
	var repeated strings.Builder
	if count == 0 {
		for range 5 {
			repeated.WriteString(character)
		}
	}
	for range count {
		repeated.WriteString(character)
	}
	return repeated.String()
}
