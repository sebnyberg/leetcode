package p2490circularsentence

import (
	"strings"
)

func isCircularSentence(sentence string) bool {
	fs := strings.Fields(sentence)
	n := len(fs)
	fs = append(fs, fs[0])
	for i := 0; i < n; i++ {
		m := len(fs[i])
		if fs[i][m-1] != fs[i+1][0] {
			return false
		}
	}
	return true
}
