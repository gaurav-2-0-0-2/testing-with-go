package iteration

import "strings"

func Repeat(char string, rc int) string {
	var repeated strings.Builder
	for i := 0; i < rc; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
