package p2103ringsandrods

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPoints(t *testing.T) {
	for _, tc := range []struct {
		rings string
		want  int
	}{
		{"B0B6G0R6R0R6G9", 1},
		{"B0R0G0R9R0B0G0", 1},
		{"G4", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rings), func(t *testing.T) {
			require.Equal(t, tc.want, countPoints(tc.rings))
		})
	}
}

func countPoints(rings string) int {
	var hasRings [10]int
	for i := 0; i < len(rings); i += 2 {
		col := rings[i]
		pos := rings[i+1] - '0'
		switch col {
		case 'B':
			hasRings[pos] |= 1
		case 'R':
			hasRings[pos] |= 2
		case 'G':
			hasRings[pos] |= 4
		}
	}

	var count int
	for _, has := range hasRings {
		if has == 7 {
			count++
		}
	}
	return count
}
