package p1090largestvaluesfromlabels

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})
	var k int
	useCount := make(map[int]int)
	var res int
	for i := 0; k < numWanted && i < len(values); i++ {
		v := values[idx[i]]
		l := labels[idx[i]]
		if useCount[l] >= useLimit {
			continue
		}
		res += v
		useCount[l]++
		k++
	}
	return res
}
