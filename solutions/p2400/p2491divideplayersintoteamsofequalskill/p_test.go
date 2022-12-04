package p2491divideplayersintoteamsofequalskill

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	n := len(skill)
	want := skill[0] + skill[n-1]
	var res int64
	for i := 0; i < n/2; i++ {
		r := n - i - 1
		if skill[i]+skill[r] != want {
			return -1
		}
		res += int64(skill[i] * skill[r])
	}
	return res
}
