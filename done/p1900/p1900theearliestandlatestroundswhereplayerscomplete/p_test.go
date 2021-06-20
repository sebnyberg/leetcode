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

	mem := make(map[uint32]struct{})
	f.tryAll(mem, state, uint8(firstPlayer), uint8(secondPlayer), 1, uint8(n))
	return []int{int(f.minMoves), int(f.maxMoves)}
}

type minMaxFinder struct {
	minMoves uint8
	maxMoves uint8
}

func (f *minMaxFinder) tryAll(mem map[uint32]struct{}, state uint32, first, second, round, n uint8) {
	if _, exists := mem[state]; exists {
		return
	}
	mem[state] = struct{}{}
	var l, r uint8
	l = 1
	r = n
	matchings := [][2]uint8{}
	var extraMask uint32 = 1 << first
	extraMask |= 1 << second
	for {
		for ; state&(1<<l) == 0; l++ {
		}
		for ; state&(1<<r) == 0; r-- {
		}
		if l >= r {
			break
		}
		if l == first && r == second {
			f.minMoves = min(f.minMoves, round)
			f.maxMoves = max(f.maxMoves, round)
			return
		}
		if l != first && r != first && l != second && r != second {
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
		f.tryAll(mem, nextState|extraMask, first, second, round+1, n)
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
