package p2300successfulpairsofspellsandpotions

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_successfulPairs(t *testing.T) {
	for _, tc := range []struct {
		spells  []int
		potions []int
		success int64
		want    []int
	}{
		{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
		{[]int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.spells), func(t *testing.T) {
			require.Equal(t, tc.want, successfulPairs(tc.spells, tc.potions, tc.success))
		})
	}
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i, sp := range spells {
		want := success / int64(sp)
		if int64(sp)*want != success {
			want++
		}
		j := sort.SearchInts(potions, int(want))
		if j == len(potions) {
			res[i] = 0
		}
		res[i] = len(potions) - j
	}
	return res
}
