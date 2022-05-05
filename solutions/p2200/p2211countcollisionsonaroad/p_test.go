package p2211countcollisionsonaroad

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countCollisions(t *testing.T) {
	for _, tc := range []struct {
		directions string
		want       int
	}{
		{"RRLLRRLSSRR", 7},
		{"SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR", 20},
		{"SSRSSRL", 3},
		{"RLRSLL", 5},
		{"LLRR", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.directions), func(t *testing.T) {
			require.Equal(t, tc.want, countCollisions(tc.directions))
		})
	}
}

func countCollisions(directions string) int {
	// Parse right: add to stack
	// Parse left: merge with previous in stack until there are no more, or do nothing
	// Parse stationary: merge with previous in stack until there are no more
	const (
		left       = 'L'
		stationary = 'S'
		right      = 'R'
	)

	stack := []byte{}
	var count int
	for _, dir := range directions {
		switch dir {
		case left:
			if len(stack) == 0 {
				continue
			}
			// Crash into previous car
			if stack[len(stack)-1] == stationary {
				count++
			} else {
				count += 2
			}
			stack[len(stack)-1] = stationary
			for len(stack) > 1 && stack[len(stack)-2] == right {
				stack[len(stack)-2] = stationary
				stack = stack[:len(stack)-1]
				count++
			}
		case right:
			stack = append(stack, right)
		case stationary:
			for len(stack) > 0 && stack[len(stack)-1] == right {
				count++
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, stationary)
			for len(stack) > 1 && stack[len(stack)-2] == right {
				stack[len(stack)-2] = stationary
				stack = stack[:len(stack)-1]
				count++
			}
		}
	}
	return count
}
