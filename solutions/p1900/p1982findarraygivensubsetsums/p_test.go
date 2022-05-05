package p1982findarraygivensubsetsums

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_recoverArray(t *testing.T) {
	for _, tc := range []struct {
		n    int
		sums []int
		want []int
	}{
		{3, []int{-3, -2, -1, 0, 0, 1, 2, 3}, []int{1, 2, -3}},
		{2, []int{0, 0, 0, 0}, []int{0, 0}},
		{4, []int{0, 0, 5, 5, 4, -1, 4, 9, 9, -1, 4, 3, 4, 8, 3, 8}, []int{0, -1, 4, 5}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, recoverArray(tc.n, tc.sums))
		})
	}
}

func recoverArray(n int, sums []int) []int {
	// Nasty problem...

	// Consider the largest and second largest sum in sums.
	// The largest sum must be the sum of all positive numbers in the original
	// array.
	// The second largest sum must either include or remove an item from its set.
	//
	// If the second largest == largest, then a zero value must have been
	// introduced / removed from the set.
	//
	// If second largest < largest, then the sum is caused either by adding
	// a very small negative number, or by subtracting a small positive number.
	//
	// So given the diff = largest - second largest,
	// Then a number diff or -diff must exist in the original array.
	// In either case, each sum either contains or does not contain the sum.
	//
	// Given the smallest sum in sums, then this sum either contains or does not
	// contain diff (or -diff). In either case, there must exist a parallel sum
	// which is equal to smallest sum + diff.
	//
	// This is how two sets sums are formed, one with and one without the number.
	// If the first set has a zero, then the number can be positive, otherwise,
	// it must be negative.

	res := make([]int, 0, 15)
	sort.Ints(sums)
	with := make([]int, 0, len(sums)/2)
	without := make([]int, 0, len(sums)/2)
	sumCount := make(map[int]int)
	for len(sums) > 1 {
		for _, sum := range sums {
			sumCount[sum]++
		}
		num := sums[len(sums)-1] - sums[len(sums)-2]
		var withHasZero bool
		for _, sum := range sums {
			if sumCount[sum] > 0 {
				if sum == 0 {
					withHasZero = true
				}
				with = append(with, sum)
				without = append(without, sum+num)
				sumCount[sum]--
				sumCount[sum+num]--
			}
		}

		if withHasZero {
			sums = with
			with, without = without[:0], without[len(sums)/2:len(sums)/2]
			res = append(res, num)
		} else {
			sums = without
			with, without = with[:0], with[len(sums)/2:len(sums)/2]
			res = append(res, -num)
		}

		// zero map
		for k := range sumCount {
			sumCount[k] = 0
		}
	}

	return res
}
