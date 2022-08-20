package p0824goatlatin

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func toGoatLatin(sentence string) string {
	words := strings.Fields(sentence)
	res := make([]string, len(words))
	for i, w := range words {
		ch := unicode.ToLower(rune(w[0]))
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			res[i] = fmt.Sprintf("%sma%s", w, bytes.Repeat([]byte{'a'}, i+1))
		} else {
			res[i] = fmt.Sprintf("%s%cma%s", w[1:], w[0], bytes.Repeat([]byte{'a'}, i+1))
		}
	}
	return strings.Join(res, " ")
}
