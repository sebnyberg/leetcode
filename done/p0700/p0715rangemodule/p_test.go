package p0715rangemodule

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeModule(t *testing.T) {
	const (
		actionAddRange    = 0
		actionRemoveRange = 1
		actionQueryRange  = 2
	)
	type any interface{}
	type testAction struct {
		name int
		args []any
		want []any
	}
	for i, tc := range []struct {
		actions []testAction
	}{
		{
			actions: []testAction{
				{actionAddRange, []any{10, 180}, nil},
				{actionAddRange, []any{150, 200}, nil},
				{actionAddRange, []any{250, 500}, nil},
				{actionQueryRange, []any{50, 100}, []any{true}},
				{actionQueryRange, []any{180, 300}, []any{false}},
				{actionQueryRange, []any{600, 1000}, []any{false}},
				{actionRemoveRange, []any{50, 150}, nil},
				{actionQueryRange, []any{50, 100}, []any{false}},
			},
		},
		{
			actions: []testAction{
				{actionAddRange, []any{10, 20}, nil},
				{actionRemoveRange, []any{14, 16}, nil},
				{actionQueryRange, []any{10, 14}, []any{true}},
				{actionQueryRange, []any{13, 15}, []any{false}},
				{actionQueryRange, []any{16, 17}, []any{true}},
			},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			r := Constructor()
			for j, a := range tc.actions {
				switch a.name {
				case actionAddRange:
					r.AddRange(a.args[0].(int), a.args[1].(int))
				case actionRemoveRange:
					r.RemoveRange(a.args[0].(int), a.args[1].(int))
				case actionQueryRange:
					res := r.QueryRange(a.args[0].(int), a.args[1].(int))
					require.Equal(t, a.want[0].(bool), res, "failed on action %v", j)
				}
			}
		})
	}
}

type RangeModule struct {
	tracked []int
}

func Constructor() RangeModule {
	return RangeModule{}
}

// Update the range [start,end] in tracked so that it contains the contents of
// sub.
func (this *RangeModule) update(start, end int, sub []int) {
	// There are more efficient ways of doing this, but it would also be more
	// complicated, so I went for the double-copy.

	// Truncate
	copy(this.tracked[start:], this.tracked[end:])
	this.tracked = this.tracked[:len(this.tracked)-(end-start)]

	// Add space for sub
	for range sub {
		this.tracked = append(this.tracked, 0)
	}
	// Move right-hand elements
	copy(this.tracked[start+len(sub):], this.tracked[start:])

	// Insert sub
	for i := range sub {
		this.tracked[start+i] = sub[i]
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	start := sort.SearchInts(this.tracked, left)
	end := sort.Search(len(this.tracked), func(i int) bool {
		return this.tracked[i] > right
	})
	sub := []int{}
	if start%2 == 0 {
		sub = append(sub, left)
	}
	if end%2 == 0 {
		sub = append(sub, right)
	}
	this.update(start, end, sub)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	start := sort.Search(len(this.tracked), func(i int) bool {
		return this.tracked[i] > left
	})
	end := sort.SearchInts(this.tracked, right)

	return start == end && start%2 == 1
}

func (this *RangeModule) RemoveRange(left int, right int) {
	start := sort.SearchInts(this.tracked, left)
	end := sort.Search(len(this.tracked), func(i int) bool {
		return this.tracked[i] > right
	})
	sub := []int{}
	if start%2 == 1 {
		sub = append(sub, left)
	}
	if end%2 == 1 {
		sub = append(sub, right)
	}
	this.update(start, end, sub)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
