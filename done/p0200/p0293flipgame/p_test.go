package p0293flipgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generatePossibleNextMoves(t *testing.T) {
	for _, tc := range []struct {
		currentState string
		want         []string
	}{
		{"++++", []string{"--++", "+--+", "++--"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.currentState), func(t *testing.T) {
			require.Equal(t, tc.want, generatePossibleNextMoves(tc.currentState))
		})
	}
}

func generatePossibleNextMoves(currentState string) []string {
	n := len(currentState)
	res := make([]string, 0)
	for i := 0; i < n-1; i++ {
		if currentState[i] == '-' {
			continue
		}
		if currentState[i] == currentState[i+1] {
			res = append(res, currentState[:i]+"--"+currentState[i+2:])
		}
	}
	return res
}
