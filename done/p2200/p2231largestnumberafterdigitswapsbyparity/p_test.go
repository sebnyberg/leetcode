package p2231largestnumberafterdigitswapsbyparity

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestInteger(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{1234, 3412},
		{65875, 87655},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, largestInteger(tc.num))
		})
	}
}

func largestInteger(num int) int {
	s := fmt.Sprint(num)
	n := len(s)
	nums := [2][]int{
		make([]int, 0, n),
		make([]int, 0, n),
	}
	for _, ch := range s {
		nums[ch&1] = append(nums[ch&1], int(ch))
	}
	sort.Ints(nums[0])
	sort.Ints(nums[1])
	var res []byte
	for _, ch := range s {
		res = append(res, byte(nums[ch&1][len(nums[ch&1])-1]))
		nums[ch&1] = nums[ch&1][:len(nums[ch&1])-1]
	}
	x, _ := strconv.Atoi(string(res))
	return x
}
