package p2272substringwithlargestvariance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestVariance(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"ykudzhiixwttnvtesiwnbcjmsydidttiyabbwzlfbmmycwjgzwhbtvtxyvkkjgfehaypiygpstkhakfasiloaveqzcywsiujvixcdnxpvvtobxgroznswwwipypwmdhldsoswrzyqthaqlbwragjrqwjxgmftjxqugoonxadazeoxalmccfeyqtmoxwbnphxih", 12},
		{"aababbb", 3},
		{"abcde", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, largestVariance(tc.s))
		})
	}
}

func largestVariance(s string) int {
	var res int
	queue := []int{}
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			if a == b {
				continue
			}
			queue = queue[:0]
			var count [26]int
			var pre bool
			for _, ch := range s {
				c := int(ch) - 'a'
				if c != a && c != b {
					continue
				}
				if c == a {
					count[a]++
					queue = append(queue, a)
				} else {
					pre = true
					count[b]++
					queue = append(queue, b)
				}
				for len(queue) > 0 && count[b] > count[a] {
					count[queue[0]]--
					queue = queue[1:]
				}
				if count[a] > 0 && pre {
					if count[b] == 0 {
						res = max(res, count[a]-1)
					} else {
						res = max(res, count[a]-count[b])
					}
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
