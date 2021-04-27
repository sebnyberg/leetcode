package p1781sumofbeautyofallsubstrings

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_beautySum(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"aabcb", 5},
		{"aabcbaa", 17},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, beautySum(tc.s))
		})
	}
}

func beautySum(s string) int {
	n := len(s)
	var totalBeauty int
	for i := range s {
		var charFreq [26]int
		countPerFreq := make([]int, 500)
		max := 0
		min := math.MaxInt32
		for j := i; j < n; j++ {
			ch := s[j] - 'a'
			f := charFreq[ch]
			countPerFreq[f]--
			countPerFreq[f+1]++
			charFreq[ch]++
			if f+1 > max {
				max = f + 1
			}
			if f+1 < min || (min == f && countPerFreq[min] == 0) {
				min = f + 1
			}
			if max > min {
				totalBeauty += max - min
			}
		}
	}
	return totalBeauty
}
