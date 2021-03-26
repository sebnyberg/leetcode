package p0229majorityelement2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, majorityElement(tc.nums))
		})
	}
}

func majorityElement(nums []int) []int {
	counts := make(map[int]int)
	n := len(nums)
	want := n / 3
	res := make([]int, 0)
	for _, n := range nums {
		if _, exists := counts[n]; !exists {
			counts[n] = want
		}
		counts[n]--
		if counts[n] < 0 {
			res = append(res, n)
			counts[n] = math.MaxInt32
		}
	}
	return res
}
