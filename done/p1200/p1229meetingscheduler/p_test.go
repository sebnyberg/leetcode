package p1229meetingscheduler

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minAvailableDuration(t *testing.T) {
	for _, tc := range []struct {
		slots1   [][]int
		slots2   [][]int
		duration int
		want     []int
	}{
		{[][]int{{10, 60}}, [][]int{{12, 17}, {21, 50}}, 8, []int{21, 29}},
		{[][]int{{10, 60}}, [][]int{{12, 17}, {12, 50}}, 8, []int{12, 20}},
		{[][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 8, []int{60, 68}},
		{[][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 12, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.slots1), func(t *testing.T) {
			require.Equal(t, tc.want, minAvailableDuration(tc.slots1, tc.slots2, tc.duration))
		})
	}
}

const (
	start = 0
	end   = 1
)

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	// not sure if slots are sorted but lets sort them anyway
	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][start] < slots1[j][start]
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][start] < slots2[j][start]
	})

	i, j := 0, 0
	n1, n2 := len(slots1), len(slots2)
	res := make([]int, 0)
	for i < n1 && j < n2 {
		a, b := slots1[i], slots2[j]
		switch {
		case a[start] > b[end]:
			j++
		case b[start] > a[end]:
			i++
		default:
			// potential match, check if overlap is big enough to be an alternative
			maxStart := max(a[start], b[start])
			minEnd := min(a[end], b[end])
			if minEnd-maxStart >= duration {
				return []int{maxStart, maxStart + duration}
			}

			// else,
			switch {
			case a[start] <= b[start]:
				if a[end] <= b[end] {
					i++
				} else {
					j++
				}
			default:
				if b[end] > a[end] {
					i++
				} else {
					j++
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
