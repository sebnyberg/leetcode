package p0899orderlyqueue

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_orderlyQueue(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"cba", 1, "acb"},
		{"baaca", 3, "aaabc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, orderlyQueue(tc.s, tc.k))
		})
	}
}

func orderlyQueue(s string, k int) string {
	// Trial and error made me realize that for k = 0, the best rotation wins
	if k == 1 {
		res := s
		for i := 1; i < len(s); i++ {
			if u := s[i:] + s[:i]; u < res {
				res = u
			}
		}
		return res
	}
	// And for k >= 2, we can shuffle the string any way we like
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	return string(bs)
}
