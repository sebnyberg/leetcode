package p1915numberofwonderfulsubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wonderfulSubstrings(t *testing.T) {
	for _, tc := range []struct {
		word string
		want int64
	}{
		{"aba", 4},
		{"aabb", 9},
		{"he", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, wonderfulSubstrings(tc.word))
		})
	}
}

func wonderfulSubstrings(word string) int64 {
	bs := []byte(word)
	n := len(word)
	var count [1 << 10]int
	bm := 0
	var res int
	count[0] = 1
	for i := 0; i < n; i++ {
		bm ^= 1 << (bs[i] - 'a')
		res += count[bm]
		count[bm]++
		for j := 0; j < 10; j++ {
			res += count[bm^(1<<j)]
		}
	}
	return int64(res)
}
