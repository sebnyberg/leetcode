package p0914xofakindinadeckofcards

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasGroupSizeX(t *testing.T) {
	for i, tc := range []struct {
		deck []int
		want bool
	}{
		{[]int{1}, false},
		{[]int{1, 1, 2, 2, 2, 2}, true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, hasGroupsSizeX(tc.deck))
		})
	}
}

func hasGroupsSizeX(deck []int) bool {
	count := make(map[int]int)
	for _, x := range deck {
		count[x]++
	}
	var g int
	for _, c := range count {
		if c == 0 {
			continue
		}
		if g == 0 {
			g = c
		} else {
			g = gcd(g, c)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
