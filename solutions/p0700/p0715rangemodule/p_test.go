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
				{actionAddRange, []any{6, 8}, nil},
				{actionRemoveRange, []any{7, 8}, nil},
				{actionRemoveRange, []any{8, 9}, nil},
				{actionAddRange, []any{8, 9}, nil},
				{actionRemoveRange, []any{1, 3}, nil},
				{actionAddRange, []any{1, 8}, nil},
				{actionQueryRange, []any{2, 4}, []any{true}},
				{actionQueryRange, []any{2, 9}, []any{true}},
				{actionQueryRange, []any{4, 6}, []any{true}},
			},
		},
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
	// If there are no elements to the right of end, truncate and append
	if end == len(this.tracked) {
		this.tracked = this.tracked[:start]
		this.tracked = append(this.tracked, sub...)
		return
	}

	// We must make len(sub) space and remove [start,end] at the same time.
	toRemove := end - start
	if toRemove < len(sub) { // add slots to end and copy
		this.tracked = append(this.tracked, make([]int, len(sub)-toRemove)...)
		copy(this.tracked[start+len(sub):], this.tracked[start:])
	} else if toRemove > len(sub) {
		// There is enough space in the original array
		// Move the sequence and truncate
		copy(this.tracked[start+len(sub):], this.tracked[end:])
		this.tracked = this.tracked[:len(this.tracked)-(toRemove-len(sub))]
	}

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
