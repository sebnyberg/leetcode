package p1405longesthappystring

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestDiverseString(t *testing.T) {
	for i, tc := range []struct {
		a    int
		b    int
		c    int
		want string
	}{
		{2, 4, 1, "bbaabbc"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestDiverseString(tc.a, tc.b, tc.c))
		})
	}
}

func longestDiverseString(a int, b int, c int) string {
	type charCount struct {
		char  byte
		count int
	}
	counts := []charCount{
		{'a', a},
		{'b', b},
		{'c', c},
	}
	s := func() {
		sort.Slice(counts, func(i, j int) bool {
			return counts[i].count > counts[j].count
		})
	}
	s()
	var res []byte
	for counts[0].count > 0 {
		kk := min(2, abs(counts[1].count-counts[0].count))
		for k := 1; k <= kk; k++ {
			res = append(res, counts[0].char)
		}
		counts[0].count -= kk
		if counts[1].count == 0 {
			break
		}
		res = append(res, counts[1].char)
		counts[1].count--
		s()
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
