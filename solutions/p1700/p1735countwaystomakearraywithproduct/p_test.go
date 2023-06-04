package p1735countwaystomakearraywithproduct

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func BenchmarkWaysToFillArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := rand.Intn(10000)
		if n < k {
			n, k = k, n
		}
		waysToFillArray([][]int{{n, k}})
	}
}

func Test_waysToFillArray(t *testing.T) {
	for i, tc := range []struct {
		queries [][]int
		want    []int
	}{
		{
			leetcode.ParseMatrix("[[15,2],[2,6],[2,16]]"),
			[]int{15, 4, 5},
		},
		{
			leetcode.ParseMatrix("[[1,1],[2,2],[3,3],[4,4],[5,5]]"),
			[]int{1, 2, 3, 10, 5},
		},
		{
			leetcode.ParseMatrix("[[2,6],[5,1],[73,660]]"),
			[]int{4, 1, 50734910},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, waysToFillArray(tc.queries))
		})
	}
}

func waysToFillArray(queries [][]int) []int {
	m := len(queries)
	mem := make(map[[2]int]int)
	res := make([]int, m)
	for i, q := range queries {
		n := q[0]
		k := q[1]
		factors := primeFactors(k)
		nways := 1
		for _, f := range factors {
			nways = (nways * nchoosek(mem, n+f-1, f)) % mod
		}
		res[i] = nways
	}

	return res
}

func primeFactors(x int) map[int]int {
	m := make(map[int]int)
	y := 2
	for x > 1 && y*y <= x {
		for x%y == 0 {
			m[y]++
			x /= y
		}
		y++
	}
	if x > 1 {
		m[x]++
	}
	return m
}

const mod = 1e9 + 7

func nchoosek(mem map[[2]int]int, n, k int) int {
	key := [2]int{n, k}
	if v, exists := mem[key]; exists {
		return v
	}
	if k == 0 {
		return 1
	}
	a := 1
	for x := n; x > k; x-- {
		a = (a * x) % mod
	}
	b := 1
	for x := n - k; x > 1; x-- {
		b = (b * x) % mod
	}
	res := a * modInverse(b, mod)
	mem[key] = res % mod
	return res % mod
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p := modPow(a, b/2, mod) % mod
	p = p * p % mod
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}
