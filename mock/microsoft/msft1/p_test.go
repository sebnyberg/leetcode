package msft1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	first, second := 0, 1
	for n > 1 {
		first, second = second, first+second
		n--
	}
	return second
}

func Test_distributeCandies(t *testing.T) {
	for _, tc := range []struct {
		candyType []int
		want      int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{6, 6, 6, 6}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.candyType), func(t *testing.T) {
			require.Equal(t, tc.want, distributeCandies(tc.candyType))
		})
	}
}

func distributeCandies(candyType []int) int {
	seen := make(map[int]bool)
	maxToEat := len(candyType) / 2
	var eaten int
	for _, candy := range candyType {
		if seen[candy] {
			continue
		}
		seen[candy] = true
		if eaten < maxToEat {
			eaten++
		}
	}

	return eaten
}
