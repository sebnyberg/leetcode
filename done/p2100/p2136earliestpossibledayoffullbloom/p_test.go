package p2136earliestpossibledayoffullbloom

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_earliestFullBloom(t *testing.T) {
	for _, tc := range []struct {
		plantTime []int
		growTime  []int
		want      int
	}{
		{[]int{27, 5, 24, 17, 27, 4, 23, 16, 6, 26, 13, 17, 21, 3, 9, 10, 28, 26, 4, 10, 28, 2},
			[]int{26, 9, 14, 17, 6, 14, 23, 24, 11, 6, 27, 14, 13, 1, 15, 5, 12, 15, 23, 27, 28, 12}, 348},
		{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
		{[]int{1, 2, 3, 2},
			[]int{2, 1, 2, 1}, 9},
		{[]int{1}, []int{1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.plantTime), func(t *testing.T) {
			require.Equal(t, tc.want, earliestFullBloom(tc.plantTime, tc.growTime))
		})
	}
}

type sorter struct {
	p []int
	g []int
}

func (f *sorter) Swap(i, j int) {
	f.p[i], f.p[j] = f.p[j], f.p[i]
	f.g[i], f.g[j] = f.g[j], f.g[i]
}

func (f *sorter) Less(i, j int) bool {
	if f.g[i] == f.g[j] {
		return f.p[i] > f.p[j]
	}
	return f.g[i] > f.g[j]
}

func (f *sorter) Len() int {
	return len(f.p)
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	f := &sorter{plantTime, growTime}
	sort.Sort(f)

	var day int
	var maxEnd int
	for i := range plantTime {
		day += plantTime[i]
		maxEnd = max(maxEnd, day+growTime[i])
	}
	return maxEnd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
