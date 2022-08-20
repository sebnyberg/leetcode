package p0825friendsofappropriateages

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numFriendRequests(t *testing.T) {
	for _, tc := range []struct {
		ages []int
		want int
	}{
		{[]int{16, 16}, 2},
		{[]int{16, 17, 18}, 2},
		{[]int{20, 30, 100, 110, 120}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ages), func(t *testing.T) {
			require.Equal(t, tc.want, numFriendRequests(tc.ages))
		})
	}
}

func numFriendRequests(ages []int) int {
	var ageCount [121]int
	for _, a := range ages {
		ageCount[a]++
	}
	var res int
	for a, cnt := range ageCount {
		for b := 1; b <= 120; b++ {
			if b > a || b > 100 && a < 100 {
				break
			}
			if cnt == 0 || float64(b) <= 0.5*float64(a)+7 {
				continue
			}
			if a == b {
				res += cnt * (cnt - 1)
			} else {
				res += cnt * ageCount[b]
			}
		}
	}
	return res
}
