package p2766relocatemarbles

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	pos := make(map[int]bool)
	for _, x := range nums {
		pos[x] = true
	}
	for i := range moveFrom {
		if pos[moveFrom[i]] {
			delete(pos, moveFrom[i])
			pos[moveTo[i]] = true
		}
	}
	vals := make([]int, 0, len(pos))
	for k := range pos {
		vals = append(vals, k)
	}
	sort.Ints(vals)
	return vals
}
