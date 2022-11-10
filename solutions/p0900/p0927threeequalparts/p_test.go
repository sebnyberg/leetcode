package p0927threeequalparts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeEqualParts(t *testing.T) {
	for i, tc := range []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, []int{2, 6}},
		{[]int{1, 0, 1, 0, 1}, []int{0, 3}},
		{[]int{1, 1, 0, 1, 1}, []int{-1, -1}},
		{[]int{1, 1, 0, 0, 1}, []int{0, 2}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, threeEqualParts(tc.arr))
		})
	}
}

func threeEqualParts(arr []int) []int {
	// Prefix zeroes can matter if the wish is to form three zeroes from arr,
	// otherwise not.
	//
	// Let's start with a sanity check for ones to rule out prefix zeroes and
	// uneven
	var ones int
	for _, x := range arr {
		if x == 1 {
			ones++
		}
	}
	if ones%3 != 0 {
		return []int{-1, -1}
	}
	if ones == 0 {
		return []int{0, 2}
	}

	// The solution must have one third of the ones in it, and the contents are
	// guided by the end. So we read the string we want from the end, then match
	// from the front.

	j := len(arr) - 1
	for k := ones / 3; ; {
		if arr[j] == 1 {
			k--
		}
		if k == 0 {
			break
		}
		j--
	}

	// arr[j:] now has the required string and starts with a 1
	//
	// From the first and ones/3th 1 in arr, validate contents.
	res := []int{}
	var k int
	for i := 0; i < j; {
		if arr[i] != 1 {
			i++
			continue
		}
		for m := j; m < len(arr); m++ {
			if arr[i] != arr[m] {
				return []int{-1, -1}
			}
			i++
		}
		if k == 0 {
			res = append(res, i-1)
		} else {
			res = append(res, i)
		}
		k++
	}
	return res
}
