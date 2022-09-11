package p2406divideinternalsintominimumnumberofgroups

import "sort"

func minGroups(intervals [][]int) int {
	// We actually want to figure out the time when the most intervals overlap.
	// One way to do so is to use coordinate compression then explore each
	// interval.
	m := make(map[int]int)
	for _, v := range intervals {
		m[v[0]]++
		m[v[1]+1]--
	}
	type item struct {
		x, val int
	}
	xs := make([]item, 0)
	for x, v := range m {
		xs = append(xs, item{x, v})
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].x < xs[j].x
	})
	var delta int
	var res int
	for _, v := range xs {
		delta += v.val
		if delta > res {
			res = delta
		}
	}
	return res
}
