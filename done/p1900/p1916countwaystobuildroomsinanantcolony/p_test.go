package p1916countwaystobuildroomsinanantcolony

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_waysToBuildRooms(t *testing.T) {
	for _, tc := range []struct {
		prevRoom []int
		want     int
	}{
		{[]int{-1, 0, 1}, 1},
		{[]int{-1, 0, 0, 1, 2}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.prevRoom), func(t *testing.T) {
			require.Equal(t, tc.want, waysToBuildRooms(tc.prevRoom))
		})
	}
}

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for i, room := range prevRoom {
		if i == 0 {
			continue
		}
		adj[room] = append(adj[room], i)
	}
	size := make([]int, n)
	calcSize(&size, adj, 0)

	nfact := 1
	for i := 2; i <= n; i++ {
		nfact *= i
		nfact %= mod
	}

	// Product of all sizes of subtrees
	den := 1
	for _, sz := range size {
		den *= sz
		den %= mod
	}

	// Inverse of denominator using Euler's theorem
	denominatorInv := modInverse(den, mod)

	return (nfact * denominatorInv) % mod
}

const (
	mod     = 1e9 + 7
	maxSize = 1e5 + 1
)

func calcSize(size *[]int, adj [][]int, cur int) int {
	sz := 1
	for _, nei := range adj[cur] {
		sz += calcSize(size, adj, nei)
	}
	(*size)[cur] = sz
	return sz
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p := modPow(a, b/2, mod) % mod
	p *= p
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}
