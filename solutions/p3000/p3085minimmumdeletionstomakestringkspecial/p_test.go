package p3085minimmumdeletionstomakestringkspecial

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDeletions(t *testing.T) {
	for _, tc := range []struct {
		word string
		k    int
		want int
	}{
		{"aabcaba", 0, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDeletions(tc.word, tc.k))
		})
	}
}

func minimumDeletions(word string, k int) int {
	var freq [26]int
	for _, c := range word {
		freq[c-'a']++
	}

	var freqs []int
	for _, f := range freq {
		if f > 0 {
			freqs = append(freqs, f)
		}
	}
	var maxFreq int
	for _, f := range freqs {
		maxFreq = max(maxFreq, f)
	}

	res := math.MaxInt32
	for kk := 0; kk <= maxFreq; kk++ {
		var count int
		for _, f := range freqs {
			if f > kk+k {
				count += f - (kk + k)
			}
			if f < kk {
				count += f
			}
		}
		res = min(res, count)
	}
	return res
}
