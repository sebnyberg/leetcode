package p1349maximumstudentstakingexam

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxStudents(t *testing.T) {
	for i, tc := range []struct {
		seats [][]byte
		want  int
	}{
		{
			[][]byte{
				[]byte("#.#"),
				[]byte("##."),
				[]byte(".#."),
			},
			3,
		},
		{
			[][]byte{
				[]byte("#.##.#"),
				[]byte(".####."),
				[]byte("#.##.#"),
			},
			4,
		},
		{
			[][]byte{
				[]byte(".#"),
				[]byte("##"),
				[]byte("#."),
				[]byte("##"),
				[]byte(".#"),
			},
			3,
		},
		{
			[][]byte{
				[]byte("#...#"),
				[]byte(".#.#."),
				[]byte("..#.."),
				[]byte(".#.#."),
				[]byte("#...#"),
			},
			10,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxStudents(tc.seats))
		})
	}
}

func maxStudents(seats [][]byte) int {
	// If there are no broken seats, the solution is easy to find: greedily
	// place students to maximize the number of "observable" seats that fall
	// outside the grid.
	//
	// But, there may be some configurations of broken seats that lead to the
	// greedy approach being sub-optimal.
	//
	// Let's explore DP as an alternative.
	//
	// What is the "state" of the world? It's the board.
	//
	// Is there a way to construct the state such that the entire board is not
	// necessary at a given point in time? Yes. The possible configurations for
	// a row is only dependent on the prior row and the row's broken seats.
	//
	//
	// Let's see how it pans out.
	m := len(seats)
	n := len(seats[0])
	room := make([]int, m)
	for i := range seats {
		for j, v := range seats[i] {
			if v == '#' {
				room[i] |= (1 << j)
			}
		}
	}
	mem := make(map[[2]int]int)
	res := dfs(mem, room, m-1, room[m-1], m, n)
	return res
}

type rowPerm struct {
	numStudents int
	nextBM      int
}

func genConfigurations(room []int, i, bm, n int) chan rowPerm {
	var nextBM int
	if i > 0 {
		nextBM = room[i-1]
	}
	var dfs func(ch chan<- rowPerm, count, cur, next, j, n int)
	dfs = func(ch chan<- rowPerm, count, cur, next, j, n int) {
		if j >= n {
			ch <- rowPerm{count, next}
			return
		}
		// Do nothing
		dfs(ch, count, cur, next, j+1, n)

		// Add a bit to this position (if valid)
		if cur&(1<<j) == 0 {
			cur |= (1 << (j + 1))
			next |= (1 << (j + 1))
			if j > 0 {
				next |= (1 << (j - 1))
			}
			dfs(ch, count+1, cur, next, j+1, n)
		}
	}
	ch := make(chan rowPerm)
	go func() {
		dfs(ch, 0, bm, nextBM, 0, n)
		close(ch)
	}()
	return ch
}

func dfs(mem map[[2]int]int, room []int, i, bm, m, n int) int {
	if i < 0 {
		return 0
	}
	key := [2]int{i, bm}
	if v, exists := mem[key]; exists {
		return v
	}
	// Need to find all valid configurations of students given the current
	// bit-mask. I'll try using a channel here as a generator - maybe it's good
	// enough even without memoization.
	var res int
	for r := range genConfigurations(room, i, bm, n) {
		res = max(res, r.numStudents+dfs(mem, room, i-1, r.nextBM, m, n))
	}
	mem[key] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
