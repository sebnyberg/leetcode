package p2251numberofflowersinfullbloom

import (
	"fmt"
	"leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fullBloomFlowers(t *testing.T) {
	for _, tc := range []struct {
		flowers [][]int
		persons []int
		want    []int
	}{
		{
			leetcode.ParseMatrix("[[1,6],[3,7],[9,12],[4,13]]"),
			[]int{2, 3, 7, 11},
			[]int{1, 2, 2, 2},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.flowers), func(t *testing.T) {
			require.Equal(t, tc.want, fullBloomFlowers(tc.flowers, tc.persons))
		})
	}
}

func fullBloomFlowers(flowers [][]int, persons []int) []int {
	const addFlower = 0
	const removeFlower = 1
	const personResult = 2

	type entry struct {
		typ    int
		idx    int
		resIdx int
	}

	entries := make([]entry, 0)
	for _, f := range flowers {
		entries = append(entries, entry{
			typ: addFlower,
			idx: f[0],
		})
		entries = append(entries, entry{
			typ: removeFlower,
			idx: f[1] + 1,
		})
	}
	for i, p := range persons {
		entries = append(entries, entry{
			typ:    personResult,
			idx:    p,
			resIdx: i,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].idx == entries[j].idx {
			return entries[i].typ < entries[j].typ
		}
		return entries[i].idx < entries[j].idx
	})
	res := make([]int, len(persons))
	var count int
	for _, e := range entries {
		switch e.typ {
		case addFlower:
			count++
		case removeFlower:
			count--
		case personResult:
			res[e.resIdx] = count
		}
	}
	return res
}
