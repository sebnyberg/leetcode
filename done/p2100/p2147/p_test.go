package p2147

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfWays(t *testing.T) {
	for _, tc := range []struct {
		corridor string
		want     int
	}{
		{"SSPPSPS", 3},
		{"PPSPSP", 1},
		{"S", 0},
		{"P", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.corridor), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfWays(tc.corridor))
		})
	}
}

const mod = 1e9 + 7

func numberOfWays(corridor string) int {
	// Corridor must contain number of chairs divisible by 2 to be valid
	var nchairs int
	for _, ch := range corridor {
		if ch == 'S' {
			nchairs++
		}
	}
	if nchairs == 0 || nchairs%2 != 0 {
		return 0
	}

	nways := 1

	for i := 0; i < len(corridor); {
		// Move i until two chairs are accounted for, or out of bounds
		var nchairs int
		for i < len(corridor) && nchairs < 2 {
			if corridor[i] == 'S' {
				nchairs++
			}
			i++
		}

		// Find the next seat, store position in j
		j := i
		for j < len(corridor) && corridor[j] == 'P' {
			j++
		}
		if j == len(corridor) {
			break
		}

		// Combine with number of valid prior arrangements
		nways = (nways * (j - i + 1)) % mod
		i = j
	}

	return nways
}

// func numberOfWays(corridor string) int {
// 	seats := make([]int, 0, len(corridor))
// 	for i, ch := range corridor {
// 		if ch == 'S' {
// 			seats = append(seats, i)
// 		}
// 	}
// 	n := len(seats)
// 	if n <= 1 || n%2 != 0 {
// 		return 0
// 	}
// 	ways := 1
// 	for i := 1; i < n-1; i += 2 {
// 		ways = (ways * (seats[i+1] - seats[i])) % mod
// 	}
// 	return ways
// }
