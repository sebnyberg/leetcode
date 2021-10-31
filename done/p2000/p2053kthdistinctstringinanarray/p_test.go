package p2053kthdistinctstringinanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthDistinct(t *testing.T) {
	for _, tc := range []struct {
		arr  []string
		k    int
		want string
	}{
		{[]string{"d", "b", "c", "b", "c", "a"}, 2, "a"},
		{[]string{"aaa", "aa", "a"}, 1, "aaa"},
		{[]string{"a", "b", "a"}, 3, ""},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, kthDistinct(tc.arr, tc.k))
		})
	}
}

func kthDistinct(arr []string, k int) string {
	count := make(map[string]int)
	for _, s := range arr {
		count[s]++
	}
	var n int
	for _, s := range arr {
		if count[s] == 1 {
			n++
			if n == k {
				return s
			}
		}
		delete(count, s)
	}
	return ""
}
