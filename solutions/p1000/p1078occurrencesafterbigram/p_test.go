package p1078occurrencesafterbigram

import "strings"

func findOcurrences(text string, first string, second string) []string {
	fs := strings.Fields(text)
	var res []string
	for i := 2; i < len(fs); i++ {
		if fs[i-2] == first && fs[i-1] == second {
			res = append(res, fs[i])
		}
	}
	return res
}
