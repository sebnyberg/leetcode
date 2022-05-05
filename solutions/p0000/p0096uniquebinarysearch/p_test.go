package p0096uniquebinarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numTrees(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{3, 5},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, numTrees(tc.n))
		})
	}
}

func numTrees(n int) (nconfigs int) {
	return numTreesCached(make(map[int]int), n)
}

func numTreesCached(mem map[int]int, n int) (nconfigs int) {
	if n <= 1 {
		return 1
	}
	for i := 1; i <= n; i++ {
		if _, exists := mem[i-1]; !exists {
			mem[i-1] = numTreesCached(mem, i-1)
		}
		if _, exists := mem[n-i]; !exists {
			mem[n-i] = numTreesCached(mem, n-i)
		}
		nconfigs += mem[i-1] * mem[n-i]
	}
	return nconfigs
}
