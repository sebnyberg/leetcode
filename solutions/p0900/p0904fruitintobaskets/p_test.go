package p0904fruitintobaskets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalFruit(t *testing.T) {
	for _, tc := range []struct {
		tree []int
		want int
	}{
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			require.Equal(t, tc.want, totalFruit(tc.tree))
		})
	}
}

func totalFruit(fruits []int) int {
	// When adding a fruit would introduce a third type,
	// remove fruits from the window until count is down to one
	m := make(map[int]int)
	var n int
	var l int
	var res int
	for r := range fruits {
		if m[fruits[r]] == 0 {
			for n >= 2 {
				if m[fruits[l]] == 1 {
					n--
				}
				m[fruits[l]]--
				l++
			}
			n++
		}
		m[fruits[r]]++
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
