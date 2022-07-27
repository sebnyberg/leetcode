package p2344mindeletionstomakearraydivisible

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		nums       []int
		numsDivide []int
		want       int
	}{
		{
			[]int{2, 2},
			[]int{964351116},
			0,
		},
		{
			[]int{40, 38, 18, 19, 18, 18, 16},
			[]int{430222122, 345833946, 609158196, 173124594, 25468560, 990277596, 295095510, 354571344, 931500936, 636837210},
			1,
		},
		{[]int{2, 3, 2, 4, 3}, []int{9, 6, 9, 3, 15}, 2},
		{[]int{4, 3, 6}, []int{8, 2, 6, 10}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums, tc.numsDivide))
		})
	}
}

func minOperations(nums []int, numsDivide []int) int {
	sort.Ints(nums)

	// De-duplicate numsDivide
	sort.Ints(numsDivide)
	var j int
	for i := range numsDivide {
		if numsDivide[i] == numsDivide[j] {
			continue
		}
		j++
		numsDivide[j] = numsDivide[i]
	}
	numsDivide = numsDivide[:j+1]

	// Find factors of first number in numsDivide
	factors := []int{}
	y := numsDivide[len(numsDivide)-1]
	for x := 1; x*x <= y; x++ {
		if y%x == 0 {
			if y != x {
				factors = append(factors, []int{x, y / x}...)
			} else {
				factors = append(factors, x)
			}
		}
	}

	// Then keep checking which factors exist in other numbers
	// Strip until done
	sort.Ints(factors)
	for _, x := range numsDivide {
		var i int
		for _, y := range factors {
			if x%y != 0 {
				continue
			}
			factors[i] = y
			i++
		}
		factors = factors[:i]
	}

	var i int
	for j, x := range nums {
		for i < len(factors) && factors[i] < x {
			i++
		}
		if i == len(factors) {
			return -1
		}
		if factors[i] == x {
			return j
		}
	}
	return -1
}
