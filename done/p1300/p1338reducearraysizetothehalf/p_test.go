package p1338reducearraysizetothehalf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSetSize(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{1, 9}, 1},
		{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{[]int{7, 7, 7, 7, 7, 7}, 1},
		{[]int{1000, 1000, 3, 7}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, minSetSize(tc.arr))
		})
	}
}

func minSetSize(arr []int) int {
	numCount := make([]uint32, 100001)
	countFreq := make([]uint32, 100001)
	for _, num := range arr {
		numCount[num]++
		countFreq[numCount[num]]++
		countFreq[numCount[num]-1]--
	}
	var nremoved int
	var actions int
	for freq := 100000; freq >= 1; freq-- {
		for countFreq[freq] > 0 {
			nremoved += freq
			actions += 1
			if nremoved >= len(arr)/2 {
				return actions
			}
			countFreq[freq]--
		}
	}
	panic("impossible")
}
