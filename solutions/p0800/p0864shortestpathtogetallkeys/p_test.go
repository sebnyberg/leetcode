package p0864shortestpathtogetallkeys

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPathAllKeys(t *testing.T) {
	for i, tc := range []struct {
		grid []string
		want int
	}{
		{[]string{"@.a..", "###.#", "b.A.B"}, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, shortestPathAllKeys(tc.grid))
		})
	}
}

const (
	isKey   = 1
	isLock  = 1 << 1
	isStart = 1 << 2
	isWall  = 1 << 3
)

func shortestPathAllKeys(grid []string) int {
	// Do not use this solution. I took it upon myself to write a suboptimal
	// solution (in terms of approach and complexity) so optimally that it would
	// bypass leetcode's TLE tests.
	//
	// It it much more efficient to use BFS
	is := [256]byte{
		'a': isKey,
		'b': isKey,
		'c': isKey,
		'd': isKey,
		'e': isKey,
		'f': isKey,
		'A': isLock,
		'B': isLock,
		'C': isLock,
		'D': isLock,
		'E': isLock,
		'F': isLock,
		'#': isWall,
		'@': isStart,
	}
	lockShift := [256]byte{
		'A': 1,
		'B': 1 << 1,
		'C': 1 << 2,
		'D': 1 << 3,
		'E': 1 << 4,
		'F': 1 << 5,
	}
	var keyCount int
	var mem [30][30][1 << 6]int
	var ii, jj int
	for i := range grid {
		for j := range grid[i] {
			ch := grid[i][j]
			keyCount += int(is[ch] & isKey)
			ii += (int(is[ch]&isStart) >> 2) * (i - ii)
			jj += (int(is[ch]&isStart) >> 2) * (j - jj)
		}
	}
	for i := range grid {
		for j := range grid[i] {
			ch := grid[i][j]
			for k := 0; k < (1 << keyCount); k++ {
				mem[i][j][k] = math.MaxInt32 - int(is[ch]&isWall>>3)*(math.MaxInt32+1)
			}
		}
	}
	m, n := len(grid), len(grid[0])
	explore(&mem, grid, &is, &lockShift, 0, ii, jj, 0, m, n)

	want := (1 << keyCount) - 1
	res := math.MaxInt32
	for i := range grid {
		for j := range grid[i] {
			if mem[i][j][want] > 0 {
				res = min(res, mem[i][j][want])
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func explore(
	mem *[30][30][1 << 6]int,
	grid []string,
	is *[256]byte,
	lockShift *[256]byte,
	steps, i, j, k, m, n int) {
	mem[i][j][k] = steps
	inside := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		ii := i + dir[0]
		jj := j + dir[1]
		if !inside(ii, jj) {
			continue
		}
		ch := grid[ii][jj]
		if b := lockShift[ch]; b > 0 && k&int(b) == 0 {
			// on top of a lock without a key
			continue
		}
		kk := k
		if is[ch] == isKey {
			kk |= 1 << int(ch-'a')
		}
		if mem[ii][jj][kk] <= steps+1 {
			continue
		}
		explore(mem, grid, is, lockShift, steps+1, ii, jj, kk, m, n)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
