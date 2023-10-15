package p2906constructproductmatrix

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_constructProductMatrix(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want [][]int
	}{
		{leetcode.ParseMatrix("[[12345],[2],[1]]"), leetcode.ParseMatrix("[[2],[0],[0]]")},
		{leetcode.ParseMatrix("[[1,2],[3,4]]"), leetcode.ParseMatrix("[[24,12],[8,6]]")},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, constructProductMatrix(tc.grid))
		})
	}
}

const mod = 12345

func inv(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	// Inverse grid, we want few, long rows.
	next := make([][]int, n)
	for i := range next {
		next[i] = make([]int, m)
	}
	for i := range grid {
		for j := range grid[i] {
			next[j][i] = grid[i][j]
		}
	}
	return next
}

func constructProductMatrix(grid [][]int) [][]int {
	// The main issue here is that since 12345 is not a prime, we cannot use
	// Euler's prime theorem to perform modular inverse multiplication.
	//
	m := len(grid)
	n := len(grid[0])
	var didInv bool
	// if m > n {
	// 	// Inverse grid, we want few, long rows.
	// 	grid = inv(grid)
	// 	m = len(grid)
	// 	n = len(grid[0])
	// 	didInv = true
	// }
	// Calculate the product of each row.
	rows := make([]int, m)
	for i := range grid {
		rows[i] = 1
		for j := range grid[i] {
			rows[i] = (rows[i] * grid[i][j]) % mod
		}
	}
	// Calculate the post-product of each row
	post := make([]int, m+1)
	for i := range post {
		post[i] = 1
	}
	for i := len(grid) - 2; i >= 0; i-- {
		post[i] = (post[i+1] * rows[i+1]) % mod
	}

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	// Calculate post-product of columnts in the row
	postCol := make([]int, n+1)
	preRow := 1
	for i := range grid {
		for j := range postCol {
			postCol[j] = 1
		}
		for j := len(grid[i]) - 2; j >= 0; j-- {
			postCol[j] = (postCol[j+1] * grid[i][j+1]) % mod
		}
		preCol := 1
		for j := range grid[i] {
			val := (postCol[j] * preCol) % mod
			val = (val * preRow) % mod
			val = (val * post[i]) % mod
			res[i][j] = val
			preCol = (preCol * grid[i][j]) % mod
		}
		preRow = (preRow * preCol) % mod
	}

	if didInv {
		res = inv(res)
	}
	return res
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
