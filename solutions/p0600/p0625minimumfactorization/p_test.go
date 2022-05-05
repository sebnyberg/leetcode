package p0625minimumfactorization

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestFactorization(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{48, 68},
		{15, 35},
		{18000000, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, smallestFactorization(tc.num))
		})
	}
}

func smallestFactorization(num int) int {
	if num == 1 {
		return 1
	}

	var parts []int
	for num > 1 {
		for x := 9; x >= 2; x-- {
			if num%x != 0 {
				continue
			}
			parts = append(parts, x)
			num /= x
			goto continueLoop
		}
		return 0
	continueLoop:
	}
	if len(parts) > 31 {
		return 0
	}
	sort.Ints(parts)
	var res int
	for _, p := range parts {
		res *= 10
		res += p
	}
	if res > (1<<31 - 1) {
		return 0
	}

	return res
}
