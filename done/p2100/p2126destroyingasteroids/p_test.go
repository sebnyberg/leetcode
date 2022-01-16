package p2126destroyingasteroids

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_asteroidsDestroyed(t *testing.T) {
	for _, tc := range []struct {
		mass      int
		asteroids []int
		want      bool
	}{
		{10, []int{3, 9, 19, 5, 21}, true},
		{5, []int{4, 9, 23, 4}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mass), func(t *testing.T) {
			require.Equal(t, tc.want, asteroidsDestroyed(tc.mass, tc.asteroids))
		})
	}
}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, a := range asteroids {
		if mass < a {
			return false
		}
		mass += a
	}
	return true
}
