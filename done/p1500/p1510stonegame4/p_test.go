package p1510stonegame4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_winnerSquareGame(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{4, true},
		{7, false},
		{17, false},
		{1, true},
		{2, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, winnerSquareGame(tc.n))
		})
	}
}

type outcome int

const (
	undecided outcome = 0
	win       outcome = 1
	loss      outcome = 2
)

func winnerSquareGame(n int) bool {
	// All perfect squares are wins
	sieve := make([]outcome, n+1)

	// Fill sieve.
	m := 0
	for {
		for ; sieve[m] != undecided; m++ {
		}
		if m == n {
			return false
		}
		sieve[m] = loss
		for k := 1; m+k*k <= n; k++ {
			idx := m + k*k
			if idx == n {
				return true
			}
			sieve[m+k*k] = win
		}
	}
}
