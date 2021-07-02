package p0974subarraysumsdivisiblebyk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subarraysDivByK(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		K    int
		want int
	}{
		{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, subarraysDivByK(tc.A, tc.K))
		})
	}
}

func subarraysDivByK(A []int, K int) int {
	sumCount := make([]uint16, K)
	sumCount[0] = 1
	var sum int
	var res int
	for _, el := range A {
		sum += el
		sum %= K
		if sum < 0 {
			sum += K
		}
		res += int(sumCount[sum])
		sumCount[sum]++
	}
	return res
}
