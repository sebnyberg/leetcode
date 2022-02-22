package p0484findpermutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPermutation(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []int
	}{
		{"I", []int{1, 2}},
		{"DI", []int{2, 1, 3}},
		{"DDDIIDDDI", []int{4, 3, 2, 1, 5, 9, 8, 7, 6, 10}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findPermutation(tc.s))
		})
	}
}

func findPermutation(s string) []int {
	// The smallest valid increase is equal to the number of decreases in a row
	// after the decrease
	res := make([]int, len(s)+1)
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 'D' {
			count = 0
			continue
		}
		count++
		res[i] = count
	}
	res[0] += 1
	maxVal := res[0]
	for i, ch := range s {
		if ch == 'D' {
			if res[i] > maxVal {
				maxVal = res[i]
			}
			res[i+1] = res[i] - 1
		} else {
			maxVal++
			res[i+1] += maxVal
		}
	}

	return res
}
