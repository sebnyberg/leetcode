package p1979findgreatestcommondivisorofarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findGCD(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 5, 6, 9, 10}, 2},
		{[]int{7, 5, 6, 8, 3}, 1},
		{[]int{3, 3}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findGCD(tc.nums))
		})
	}
}

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return gcd(min, max)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
