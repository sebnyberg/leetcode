package p2288applydiscounttoprices

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Fields(sentence)

	isNum := func(s string) bool {
		if len(s) <= 1 {
			return false
		}
		if s[0] != '$' {
			return false
		}
		for i := 1; i < len(s); i++ {
			if !unicode.IsNumber(rune(s[i])) {
				return false
			}
		}
		return true
	}

	frac := float64(100-discount) / 100
	for i, w := range words {
		if !isNum(w) {
			continue
		}
		x, err := strconv.Atoi(w[1:])
		if err != nil {
			panic(err)
		}
		newPrice := float64(x) * frac
		words[i] = fmt.Sprintf("$%.02f", newPrice)
	}
	return strings.Join(words, " ")
}
