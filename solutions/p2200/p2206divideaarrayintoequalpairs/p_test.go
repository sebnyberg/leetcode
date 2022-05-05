package p2206

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divideArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{3, 2, 3, 2, 2, 2}, true},
		{[]int{1, 2, 3, 4}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, divideArray(tc.nums))
		})
	}
}

func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for _, count := range m {
		if count%2 != 0 {
			return false
		}
	}
	return true
}
