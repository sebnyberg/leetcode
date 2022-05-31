package p0735asteroidcollision

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_asteroidCollision(t *testing.T) {
	for _, tc := range []struct {
		asteroids []int
		want      []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.asteroids), func(t *testing.T) {
			require.Equal(t, tc.want, asteroidCollision(tc.asteroids))
		})
	}
}

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, a := range asteroids {
		if a > 0 {
			// Going right
			stack = append(stack, a)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -a {
			// Remove last element from stack
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, a)
			continue
		}
		if stack[len(stack)-1] == -a {
			stack = stack[:len(stack)-1]
		} else {
			// Do not add anything, this meteor just exploded.
		}
	}
	return stack
}
