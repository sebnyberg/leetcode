package p1663smalleststr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSmallestString(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want string
	}{
		// {3, 27, "aay"},
		{5, 73, "aaszz"},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.n, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, getSmallestString(tc.n, tc.k))
		})
	}
}

func getSmallestString(n int, k int) string {
	res := make([]rune, n)
	cur := k - n
	if k > 26*n {
		return ""
	}
	for i := n - 1; i >= 0; i-- {
		if cur <= 25 {
			res[i] = rune('a' + cur)
			cur = 0
			continue
		}
		res[i] = 'z'
		cur -= 25
	}
	return string(res)
}
