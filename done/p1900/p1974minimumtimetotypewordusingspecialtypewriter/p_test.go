package p1974minimumtimetotypewordusingspecialtypewriter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minTimeToType(t *testing.T) {
	for _, tc := range []struct {
		word string
		want int
	}{
		{"abc", 5},
		{"bza", 7},
		{"zjpc", 34},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, minTimeToType(tc.word))
		})
	}
}

func minTimeToType(word string) int {
	var res int
	cur := 0
	for _, ch := range word {
		res++
		val := int(ch - 'a')
		dist := abs(cur - val)
		if val < cur {
			dist = min(dist, val+26-cur)
		} else {
			dist = min(dist, cur+26-val)
		}
		res += dist
		cur = val
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
