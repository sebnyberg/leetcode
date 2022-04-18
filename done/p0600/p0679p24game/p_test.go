package p0679p24game

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_judgePoint24(t *testing.T) {
	for _, tc := range []struct {
		cards []int
		want  bool
	}{
		{[]int{1, 9, 1, 2}, true},
		{[]int{8, 1, 2}, true},
		{[]int{4, 1, 8, 7}, true},
		{[]int{1, 2, 1, 2}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.cards), func(t *testing.T) {
			require.Equal(t, tc.want, judgePoint24(tc.cards))
		})
	}
}

func judgePoint24(cards []int) bool {
	// The idea is to work with bitmaps, and possible values.
	// Each bitmap, e.g. 1000, 1011, etc has a list of possible values, stored in
	// vals.
	// Once we are done, we explore whether there exists a val ~= 24 for the
	// bitmap containing all cards.
	var vals [1 << 4]map[float64]struct{}

	// The base case is bitmaps containing a single card, these can only have
	// themselves as a value
	bms := []int{}
	for i := 0; i < len(cards); i++ {
		vals[1<<i] = map[float64]struct{}{
			float64(cards[i]): {},
		}
		bms = append(bms, 1<<i)
	}

	// For each "current" bitmask which is not equal to the final bitmask, combine
	// it with all prior bitmasks such that they do not have a card in common. Add
	// the result to the next round.
	next := []int{}
	curr := make([]int, len(bms))
	copy(curr, bms)
	allCards := (1 << len(cards)) - 1
	for len(curr) > 0 {
		next = next[:0]
		for _, bm1 := range curr {
			if bm1 == allCards {
				continue
			}
			for _, bm2 := range bms {
				if bm1&bm2 > 0 { // there must be no overlap
					continue
				}
				// bm & otherBM is a valid combination. Add it to next, and add all
				// possible values that could come out of combining the cards found in
				comb := bm1 | bm2
				if vals[comb] == nil {
					vals[comb] = make(map[float64]struct{})
				}
				for v1 := range vals[bm1] {
					for v2 := range vals[bm2] {
						if v2 != 0 {
							vals[comb][v1/v2] = struct{}{}
						}
						if v1 != 0 {
							vals[comb][v2/v1] = struct{}{}
						}
						vals[comb][v1+v2] = struct{}{}
						vals[comb][v1-v2] = struct{}{}
						vals[comb][v2-v1] = struct{}{}
						vals[comb][v1*v2] = struct{}{}
					}
				}
				next = append(next, comb)
				bms = append(bms, comb)
			}
		}

		curr, next = next, curr
	}

	const eps = 0.001

	for v := range vals[allCards] {
		if math.Abs(v-24) < eps {
			return true
		}
	}

	return false
}
