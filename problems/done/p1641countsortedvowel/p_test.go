package p1641countsortedvowel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVowelStrings(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1, 5},
		{2, 15},
		{33, 66045},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countVowelStrings(tc.n))
		})
	}
}

func sum(ns []int) (res int) {
	for _, n := range ns {
		res += n
	}
	return res
}

func countVowelStrings(n int) int {
	ns := make([]int, 5)
	for i := range ns {
		ns[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := len(ns) - 1; j > 0; j-- {
			ns[j] = sum(ns[:j+1])
		}
	}
	return sum(ns)
}
