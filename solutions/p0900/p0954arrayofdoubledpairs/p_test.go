package p0954arrayofdoubledpairs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canReorderDoubled(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want bool
	}{
		{[]int{3, 1, 3, 6}, false},
		{[]int{2, 1, 2, 6}, false},
		{[]int{4, -2, 2, -4}, true},
		{[]int{4, -2, 2, -4, 1}, true},
		{[]int{1, 2, 4, 16, 8, 4}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, canReorderDoubled(tc.arr))
		})
	}
}

func canReorderDoubled(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < 0 && arr[j] < 0 {
			return arr[i] > arr[j]
		}
		return arr[i] < arr[j]
	})
	numCount := make(map[int]uint16, len(arr))
	for _, num := range arr {
		numCount[num]++
	}
	var faults int
	for _, num := range arr {
		if numCount[num] == 0 {
			continue
		}
		numCount[num]--
		if numCount[num*2] > 0 {
			numCount[num*2]--
			continue
		}
		faults++
		if faults > 1 {
			break
		}
	}
	if len(arr)%2 == 1 {
		return faults <= 1
	}
	return faults == 0
}
