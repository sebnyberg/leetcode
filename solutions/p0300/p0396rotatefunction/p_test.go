package p0396rotatefunction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxRotateFunction(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{4, 3, 2, 6}, 26},
		{[]int{1000000007}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxRotateFunction(tc.nums))
		})
	}
}

func maxRotateFunction(nums []int) int {
	n := len(nums)
	var sum int
	var prod int
	for i, num := range nums {
		sum += num
		prod += i * num
	}
	// Rotating is equal to adding all numbers to the product and removing
	// a certain index from the list
	maxProd := prod
	for i := n - 1; i > 0; i-- {
		prod += sum
		prod -= nums[i] * n
		if prod > maxProd {
			maxProd = prod
		}
	}
	return maxProd
}
