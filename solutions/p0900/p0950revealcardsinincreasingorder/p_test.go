package p0950revealcardsinincreasingorder

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	for i, tc := range []struct {
		deck []int
		want []int
	}{
		{[]int{17, 13, 11, 2, 3, 5, 7}, []int{2, 13, 3, 11, 5, 17, 7}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, deckRevealedIncreasing(tc.deck))
		})
	}
}

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	placed := make([]bool, n)
	var j int
	k := 1
	sort.Ints(deck)
	res := make([]int, n)
	placed[0] = true
	res[0] = deck[0]
	for x := 0; k < len(deck); x++ {
		for placed[j] {
			j = (j + 1) % n
		}
		if x&1 == 1 {
			placed[j] = true
			res[j] = deck[k]
			k++
		}
		j = (j + 1) % n
	}
	return res
}
