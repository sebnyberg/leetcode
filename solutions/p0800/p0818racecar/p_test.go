package p0818racecar

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_racecar(t *testing.T) {
	for _, tc := range []struct {
		target int
		want   int
	}{
		{4, 5},
		{3, 2},
		{6, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, racecar(tc.target))
		})
	}
}

func racecar(target int) int {
	type state struct {
		position int
		speed    int
	}
	curr := []state{{0, 1}}
	// Perform BFS, avoiding ending up in the same state more than once, and
	// really silly values such as target * 2
	next := []state{}
	stateSeen := make(map[state]struct{})
	seen := func(s state) bool {
		_, exists := stateSeen[s]
		return exists
	}
	for steps := 0; ; steps++ {
		next = next[:0]
		for _, x := range curr {
			if x.position == target {
				return steps
			}
			// Accelerate
			a := state{x.position + x.speed, x.speed * 2}
			if a.position > 0 && a.position < target*2 && !seen(a) {
				stateSeen[a] = struct{}{}
				next = append(next, a)
			}
			r := state{x.position, -1}
			if x.speed < 0 {
				r.speed = 1
			}
			if !seen(r) {
				stateSeen[r] = struct{}{}
				next = append(next, r)
			}
		}
		curr, next = next, curr
	}
}
