package p0681nextclosesttime

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextClosestTime(t *testing.T) {
	for _, tc := range []struct {
		time string
		want string
	}{
		{"23:59", "22:22"},
		{"19:34", "19:39"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.time), func(t *testing.T) {
			require.Equal(t, tc.want, nextClosestTime(tc.time))
		})
	}
}

func nextClosestTime(time string) string {
	timeBytes := []byte(time)
	nums := []byte{}
	for _, b := range timeBytes {
		if b == ':' {
			continue
		}
		nums = append(nums, b)
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// First option, if the last minute can be increased, it is the answer
	for _, b := range nums {
		if b > timeBytes[4] {
			timeBytes[4] = b
			return string(timeBytes)
		}
	}

	// Second option, set timeBytes[4] to the lowest possible value,
	// then find the first non-4 index to increment.
	timeBytes[4] = nums[0]
	for _, b := range nums {
		if b > timeBytes[3] && b < '6' {
			timeBytes[3] = b
			return string(timeBytes)
		}
	}
	timeBytes[3] = nums[0]
	for _, b := range nums {
		if b > timeBytes[1] {
			if b > '3' && timeBytes[0] == '2' {
				break
			}
			timeBytes[1] = b
			return string(timeBytes)
		}
	}
	timeBytes[1] = nums[0]
	for _, b := range nums {
		if b > timeBytes[0] && b < '2' || (b == '2' && timeBytes[1] < '4') {
			timeBytes[0] = b
			return string(timeBytes)
		}
	}

	timeBytes[0] = nums[0]

	return string(timeBytes)
}

func timeValid(t []byte) bool {
	if t[3] > 5 {
		return false
	}
	if t[0] > 2 || t[0] == 1 && t[1] > 3 {
		return false
	}
	return true
}
