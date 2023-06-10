package p1408stringmatchinginanarray

import "sort"

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var res []string
outer:
	for i := 0; i < len(words); i++ {
		m := len(words[i])
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			n := len(words[j])
			for k := 0; k+m <= n; k++ {
				if words[j][k:k+m] == words[i] {
					res = append(res, words[i])
					continue outer
				}
			}
		}
	}
	return res
}
