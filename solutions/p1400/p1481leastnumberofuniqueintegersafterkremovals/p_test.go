package p1481leastnumberofuniqueintegersafterkremovals

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	type a struct {
		x   int
		cnt int
	}
	var counts []a
	m := make(map[int]int)
	for _, x := range arr {
		if _, exists := m[x]; !exists {
			m[x] = len(counts)
			counts = append(counts, a{x: x, cnt: 1})
		} else {
			counts[m[x]].cnt++
		}
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].cnt < counts[j].cnt
	})
	var j int
	for j < len(counts) && k-counts[j].cnt >= 0 {
		k -= counts[j].cnt
		j++
	}
	return len(counts) - j
}
