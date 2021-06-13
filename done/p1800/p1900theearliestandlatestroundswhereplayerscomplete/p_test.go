package p1900theearliestandlatestroundswhereplayerscomplete

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_earliestAndLatest(t *testing.T) {
	for _, tc := range []struct {
		n, firstPlayer, secondPlayer int
		want                         []int
	}{
		// {11, 2, 4, []int{3, 4}},
		// {5, 1, 5, []int{1, 1}},
		{4, 3, 4, []int{2, 2}},
		{28, 1, 12, []int{2, 2}},
	} {
		t.Run(fmt.Sprintf("%v/%v/%v", tc.n, tc.firstPlayer, tc.secondPlayer), func(t *testing.T) {
			require.Equal(t, tc.want, earliestAndLatest(tc.n, tc.firstPlayer, tc.secondPlayer))
		})
	}
}

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	f := minMaxFinder{
		minMoves: 30,
		maxMoves: 1,
	}

	var state uint32
	var bit uint32 = 1 << 1
	for i := 1; i <= n; i++ {
		state |= bit
		bit <<= 1
	}

	f.tryAll(state, uint8(firstPlayer), uint8(secondPlayer), 1, uint8(n))
	return []int{int(f.minMoves), int(f.maxMoves)}
}

type minMaxFinder struct {
	minMoves uint8
	maxMoves uint8
}

func (f *minMaxFinder) tryAll(state uint32, first, second, round, n uint8) {
	// fmt.Printf("round: %v\t bitmask: %v\n", round, debugBitMask(state, n))
	// Try all possible matchings inside the state
	// If a matching is first vs second, update both min and max moves
	var l, r uint8
	l = 1
	r = n
	matchings := [][2]uint8{}
	var extraMask uint32
	for {
		for ; state&(1<<l) == 0; l++ {
		}
		for ; state&(1<<r) == 0; r-- {
		}
		if l >= r {
			break
		}
		switch {
		case l == first && r == second:
			f.minMoves = min(f.minMoves, round)
			f.maxMoves = max(f.maxMoves, round)
			return
		case l == first || l == second:
			extraMask |= 1 << l
		case r == first || r == second:
			extraMask |= 1 << r
		default:
			matchings = append(matchings, [2]uint8{l, r})
		}
		l++
		r--
	}
	// Add middle player
	if l == r {
		extraMask |= 1 << l
	}
	nextStates := make(chan uint32, 3)
	go func() {
		genNextStates(nextStates, matchings, 0, 0, uint8(len(matchings)))
		close(nextStates)
	}()
	for nextState := range nextStates {
		// fmt.Printf("round: %v\t extra  : %v\n", round, debugBitMask(extraMask, n))
		f.tryAll(nextState|extraMask, first, second, round+1, n)
	}
}

func genNextStates(nextState chan uint32, matchings [][2]uint8, state uint32, idx, n uint8) {
	if idx == n {
		nextState <- state
		return
	}
	genNextStates(nextState, matchings, state|1<<matchings[idx][0], idx+1, n)
	genNextStates(nextState, matchings, state|1<<matchings[idx][1], idx+1, n)
}

func max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

func debugBitMask(bm uint32, n uint8) string {
	b := uint32(2)
	out := make([]byte, 0, n)
	for i := uint8(0); i < n; i++ {
		if bm&b > 0 {
			out = append(out, 'X')
		} else {
			out = append(out, '-')
		}
		b <<= 1
	}
	return string(out)
}
