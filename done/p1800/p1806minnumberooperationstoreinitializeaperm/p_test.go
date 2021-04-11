package p1806minnumberooperationstoreinitializeaperm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reinitializePermutation(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 1},
		{4, 2},
		{6, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, reinitializePermutation(tc.n))
		})
	}
}

func reinitializePermutation(n int) int {
	// Assume that there are no cycles
	// Assume that there is never more than 1000 operations

	var perms [2][]int
	perms[0] = make([]int, n)
	perms[1] = make([]int, n)
	for i := 0; i < n; i++ {
		perms[0][i] = i
	}

	iter := 1
	for {
		permIdx := iter % 2
		prev := (iter - 1) % 2
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				perms[permIdx][i] = perms[prev][i/2]
			} else {
				perms[permIdx][i] = perms[prev][n/2+(i-1)/2]
			}
		}
		if check(perms[permIdx]) {
			return iter
		}
		iter++
	}
}

func check(perms []int) bool {
	for i, n := range perms {
		if n != i {
			return false
		}
	}
	return true
}
