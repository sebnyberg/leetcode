package p3016minimumnumberofpushestotypewordii

import "sort"

func minimumPushes(word string) int {
	counts := make([]int, 26)
	for _, ch := range word {
		counts[ch-'a']++
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	var res int
	for i, cc := range counts {
		res += ((i / 8) + 1) * cc
	}
	return res
}
