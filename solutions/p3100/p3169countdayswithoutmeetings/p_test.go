package p3169countdayswithoutmeetings

import "sort"

func countDays(days int, meetings [][]int) int {
	type change struct {
		t int
		d int
	}
	delta := make([]change, len(meetings)*2)
	for i := range meetings {
		delta[i*2] = change{meetings[i][0], 1}
		delta[i*2+1] = change{meetings[i][1] + 1, -1}
	}
	sort.Slice(delta, func(i, j int) bool {
		return delta[i].t < delta[j].t
	})
	var curr int
	var j int
	var res int
	for day := 1; day <= days; {
		for j < len(delta) && delta[j].t <= day {
			curr += delta[j].d
			j++
		}
		next := days + 1
		if j < len(delta) {
			next = delta[j].t
		}
		if curr == 0 {
			res += next - day
		}
		day = next
	}
	return res
}
