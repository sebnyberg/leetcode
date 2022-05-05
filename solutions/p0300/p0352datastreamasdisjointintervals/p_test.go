package p0352datastreamasdisjointintervals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSummaryRanges(t *testing.T) {
	// t.Run("first", func(t *testing.T) {
	// 	sum := Constructor()
	// 	sum.AddNum(1)
	// 	res := sum.GetIntervals()
	// 	require.Equal(t, [][]int{{1, 1}}, res)
	// 	sum.AddNum(3)
	// 	res = sum.GetIntervals()
	// 	require.Equal(t, [][]int{{1, 1}, {3, 3}}, res)
	// 	sum.AddNum(7)
	// 	res = sum.GetIntervals()
	// 	require.Equal(t, [][]int{{1, 1}, {3, 3}, {7, 7}}, res)
	// 	sum.AddNum(2)
	// 	res = sum.GetIntervals()
	// 	require.Equal(t, [][]int{{1, 3}, {7, 7}}, res)
	// 	sum.AddNum(6)
	// 	res = sum.GetIntervals()
	// 	require.Equal(t, [][]int{{1, 3}, {6, 7}}, res)
	// })

	t.Run("second", func(t *testing.T) {
		sum := Constructor()
		for _, tc := range []struct {
			in   int
			want [][]int
		}{
			{6, [][]int{{6, 6}}},
			{6, [][]int{{6, 6}}},
			{0, [][]int{{0, 0}, {6, 6}}},
			{4, [][]int{{0, 0}, {4, 4}, {6, 6}}},
			{8, [][]int{{0, 0}, {4, 4}, {6, 6}, {8, 8}}},
			{7, [][]int{{0, 0}, {4, 4}, {6, 8}}},
			{6, [][]int{{0, 0}, {4, 4}, {6, 8}}},
			{4, [][]int{{0, 0}, {4, 4}, {6, 8}}},
			{7, [][]int{{0, 0}, {4, 4}, {6, 8}}},
			{5, [][]int{{0, 0}, {4, 8}}},
		} {
			t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
				sum.AddNum(tc.in)
				res := sum.GetIntervals()
				require.Equal(t, tc.want, res)
			})
		}
	})
}

type SummaryRanges struct {
	intervals *dsuInterval
	exists    []bool
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{
		intervals: newDSUInterval(10000),
		exists:    make([]bool, 10002),
	}
}

func (this *SummaryRanges) AddNum(val int) {
	if this.exists[val] { // noop
		return
	}
	this.exists[val] = true
	if val > 0 && this.exists[val-1] {
		this.intervals.union(val-1, val)
	}
	if this.exists[val+1] {
		this.intervals.union(val, val+1)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, 0)
	for i := 0; i < 10001; i++ {
		if this.exists[i] {
			end := this.intervals.maxVal[i]
			res = append(res, []int{i, this.intervals.maxVal[i]})
			i += end - i
		}
	}
	return res
}

type dsuInterval struct {
	parent []int
	maxVal []int
}

func newDSUInterval(n int) *dsuInterval {
	parent := make([]int, n+1)
	maxVal := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
		maxVal[i] = i
	}
	return &dsuInterval{parent, maxVal}
}

// find returns the root index for the provided number
func (d *dsuInterval) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	rootIdx := d.find(d.parent[a])
	if rootIdx != a {
		d.parent[a] = rootIdx // Path compression
	}
	return rootIdx
}

func (d *dsuInterval) union(a, b int) {
	ra, rb := d.find(a), d.find(b)
	if ra != rb {
		if rb < ra {
			ra, rb = rb, ra
		}
		d.parent[rb] = ra
		d.maxVal[ra] = max(d.maxVal[ra], d.maxVal[rb])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
