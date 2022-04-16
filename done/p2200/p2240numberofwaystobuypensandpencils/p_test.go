package p2240

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_waysToBuyPensPencils(t *testing.T) {
	for _, tc := range []struct {
		total        int
		cost1, cost2 int
		want         int64
	}{
		{20, 10, 5, 9},
		{5, 10, 10, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.total), func(t *testing.T) {
			require.Equal(t, tc.want, waysToBuyPensPencils(tc.total, tc.cost1, tc.cost2))
		})
	}
}

// func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
// 	// If you buy zero of the first, you can buy total/cost2 of the second
// 	// If you buy one of the first, you can buy (total-cost1)/cost2 of the second

// 	if cost1 <= cost2 {
// 		cost1, cost2 = cost2, cost1
// 	}

// 	// Add cost2 until cost2%cost1 == 0
// 	counts := []int{1}
// 	curr := cost1
// 	for {
// 		idx := (curr - 1) / cost1
// 		if idx >= len(counts) {
// 			counts = append(counts, 0)
// 		}
// 		counts[idx]++
// 		if curr%cost1 == 0 {
// 			break
// 		}
// 	}
// 	res := 1
// 	res += total / cost1
// 	res += total / cost2
// 	curr = 0
// 	for i := 0; i < len(counts); i++ {
// 		curr += cost1
// 		if curr+cost2 >= total {
// 			break
// 		}
// 		repeatsEvery := len(counts) * cost1
// 	}
// 	return 0
// }

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	// Brute-force will work, because in the extreme cases, the execution will
	// still be fast enough.
	if cost1 < cost2 {
		cost1, cost2 = cost2, cost1
	}

	var res int
	for x := 0; x <= total; x += cost1 {
		d := (total - x) / cost2
		d++
		res += d
	}

	return int64(res)
}
