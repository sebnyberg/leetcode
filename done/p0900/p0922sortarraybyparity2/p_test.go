package p0922sortarraybyparity2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortArrayByParityII(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{[]int{2, 3}, []int{2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, sortArrayByParityII(tc.nums))
		})
	}
}

func sortArrayByParityII(nums []int) []int {
	var i, j int
	res := make([]int, len(nums))
	for _, num := range nums {
		if num%2 == 0 {
			res[i*2] = num
			i++
		} else {
			res[j*2+1] = num
			j++
		}
	}
	return res
}
