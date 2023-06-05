package p1751maximumnumberofeventsthatcanbeattendedii

import (
	"math"
	"sort"
)

func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		a := events[i]
		b := events[j]
		return a[0] < b[0]
	})
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, k+1)
		for j := range mem[i] {
			mem[i][j] = math.MinInt32
		}
	}
	res := dp(mem, events, 0, k)
	return res
}

func dp(mem, events [][]int, i, k int) int {
	if i == len(events) {
		return 0
	}
	if k == 0 {
		return 0
	}
	if mem[i][k] != math.MinInt32 {
		return mem[i][k]
	}
	e := events[i]
	nextEvent := sort.Search(len(events[i+1:]), func(j int) bool {
		return events[i+1+j][0] > e[1]
	}) + i + 1
	res := max(
		dp(mem, events, i+1, k),
		events[i][2]+dp(mem, events, nextEvent, k-1),
	)
	mem[i][k] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
