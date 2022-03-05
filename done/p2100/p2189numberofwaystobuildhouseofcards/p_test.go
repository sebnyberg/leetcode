package p2189numberofwaystobuildhouseofcards

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_houseOfCards(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{16, 2},
		{2, 1},
		{4, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, houseOfCards(tc.n))
		})
	}
}

func houseOfCards(n int) int {
	var baseCardWays [200][501]int

	// Now let's consider a base-k house. Such a house has k*2+(k-1) cards
	// On top of that house, we may build any base-m house where m < k
	// If we calculate the number of ways we can build a card consisting of
	// n cards for each base, then the recursion can find the solution.
	var res int
	for k := 1; k*2+k-1 <= n; k++ {
		used := k*3 - 1
		remains := n - used
		baseCardWays[k][used] = 1
		for m := k - 1; m >= 1; m-- {
			for numCards := 2; numCards <= remains; numCards++ {
				baseCardWays[k][used+numCards] += baseCardWays[m][numCards]
			}
		}
		res += baseCardWays[k][n]
	}
	return res
}
