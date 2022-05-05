package p0294flipgame2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canWin(t *testing.T) {
	for _, tc := range []struct {
		currentState string
		want         bool
	}{
		{"++++", true},
		{"+", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.currentState), func(t *testing.T) {
			require.Equal(t, tc.want, canWin(tc.currentState))
		})
	}
}

func canWin(currentState string) bool {
	return canWinRecursive(make(map[string]bool), len(currentState), currentState)
}

func canWinRecursive(mem map[string]bool, n int, state string) bool {
	if v, exists := mem[state]; exists {
		return v
	}
	for i := 0; i < n-1; i++ {
		if state[i:i+2] != "++" {
			continue
		}
		// Can make move
		nextState := state[:i] + "--" + state[i+2:]
		if otherCanWin := canWinRecursive(mem, n, nextState); !otherCanWin {
			mem[state] = true
			return true
		}
	}
	mem[state] = false
	return false
}
