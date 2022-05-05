package p2198numberofsingledivisortriplets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleDivisorTriplet(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{4, 6, 7, 3, 2}, 12},
		{[]int{1, 2, 2}, 6},
		{[]int{1, 1, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, singleDivisorTriplet(tc.nums))
		})
	}
}

func singleDivisorTriplet(nums []int) int64 {
	// Since the maximum number of nums[i] <= 100, we can count the frequency
	// of each number and find triplets that way instead...
	// The difficulty of this exercise is not so much about finding triplets but
	// rather removing the duplicate triplets when counting.
	var numCount [101]int
	isDivisible := func(x, y int) int {
		if x%y == 0 {
			return 1
		}
		return 0
	}
	for _, n := range nums {
		numCount[n]++
	}
	var res int
	for x := 1; x <= 100; x++ {
		if numCount[x] == 0 {
			continue
		}
		for y := x; y <= 100; y++ {
			if numCount[y] == 0 {
				continue
			}
			for z := y; z <= 100; z++ {
				if numCount[z] == 0 {
					continue
				}
				d := x + y + z
				nDiv := isDivisible(d, x) + isDivisible(d, y) + isDivisible(d, z)
				if nDiv != 1 {
					continue
				}
				if x == y {
					res += numCount[x] * (numCount[y] - 1) / 2 * numCount[z]
				} else if y == z {
					res += numCount[x] * numCount[y] * (numCount[z] - 1) / 2
				} else {
					res += numCount[x] * numCount[y] * numCount[z]
				}
			}
		}
	}
	return int64(6 * res)
}
