package p0937reorderdatainlogfiles

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	var digits []string
	var letters []string
	for _, l := range logs {
		p := strings.Fields(l)
		if unicode.IsLetter(rune(p[1][0])) {
			letters = append(letters, l)
		} else {
			digits = append(digits, l)
		}
	}
	sort.Slice(letters, func(i, j int) bool {
		p := strings.Fields(letters[i])
		q := strings.Fields(letters[j])
		left := strings.Join(p[1:], " ")
		right := strings.Join(q[1:], " ")
		if left == right {
			return p[0] < q[0]
		}
		return left < right
	})
	letters = append(letters, digits...)
	return letters
}
