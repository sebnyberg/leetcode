package p2055platesbetweencandles

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_platesBetweenCandles(t *testing.T) {
	for _, tc := range []struct {
		s       string
		queries [][]int
		want    []int
	}{
		{"**|**|***|", [][]int{{2, 5}, {5, 9}}, []int{2, 3}},
		{"***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}, []int{9, 0, 0, 0, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, platesBetweenCandles(tc.s, tc.queries))
		})
	}
}

func platesBetweenCandles(s string, queries [][]int) []int {
	// Looks like a typical prefix sum exercise.
	// Iterate over the input
	// When encountering plates, add to current count of plates..
	// When encountering a candle ('|'), update current count of plates between
	plateCount := make([]int, len(s))
	prevCandle := make([]int, len(s))
	prev := -1
	count := 0
	for i := range s {
		if s[i] == '|' {
			prev = i
		} else {
			count++
		}
		plateCount[i] = count
		prevCandle[i] = prev
	}
	nextCandle := make([]int, len(s))
	next := math.MaxInt32
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			next = i
		}
		nextCandle[i] = next
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		start, end := nextCandle[q[0]], prevCandle[q[1]]
		if start < end {
			res[i] = plateCount[end] - plateCount[start]
		}
	}

	return res
}
