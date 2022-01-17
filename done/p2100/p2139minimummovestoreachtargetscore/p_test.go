package p2139minimummovestoreachtargetscore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMoves(t *testing.T) {
	for _, tc := range []struct {
		target     int
		maxDoubles int
		want       int
	}{
		{5, 0, 4},
		{19, 2, 7},
		{10, 4, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, minMoves(tc.target, tc.maxDoubles))
		})
	}
}

func minMoves(target int, maxDoubles int) int {
	var moves int
	for target > 1 && maxDoubles > 0 {
		if target%2 == 1 {
			moves++
		}
		maxDoubles--
		moves++
		target /= 2
	}
	for target > 1 {
		moves++
		target--
	}
	return moves
}
