package p1728catandmouseii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canMouseWin(t *testing.T) {
	for i, tc := range []struct {
		grid      []string
		catJump   int
		mouseJump int
		want      bool
	}{
		{
			[]string{
				".....M",
				"##.#.#",
				"#....#",
				".##...",
				"...C..",
				"F.....",
				".#..#.",
			},
			1, 5, false,
		},
		{
			[]string{
				"###.#",
				"##.F#",
				"C###.",
				".#M.#",
				"####.",
			},
			2, 2, false,
		},
		{
			[]string{
				"..#C",
				"...#",
				"..M.",
				"#F..",
				"....",
			},
			2, 1, true,
		},
		{
			[]string{"####F", "#C...", "M...."},
			1, 2, true,
		},
		{
			[]string{"MC..F"},
			1, 3, true,
		},
		{
			[]string{"M.C...F"},
			1, 3, false,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, canMouseWin(tc.grid, tc.catJump, tc.mouseJump))
		})
	}
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

	const (
		mouse        = 0
		cat          = 1
		uncertain    = 1 << 0
		mouseWillWin = 1 << 1
		catWillWin   = 1 << 2
	)

	m := len(grid)
	n := len(grid[0])

	idx := func(i, j int) int {
		return i*n + j
	}
	pos := func(ij int) (int, int) {
		return ij / n, ij % n
	}

	// res[m][c][mouse] = mouse at mi, cat at ci, mouse's turn, value = outcome
	res := make([][][2]int, m*n)

	// Mark all locations as uncertain
	for mi := range res {
		res[mi] = make([][2]int, m*n)
		for ci := range res[mi] {
			res[mi][ci][mouse] = uncertain
			res[mi][ci][cat] = uncertain
		}
	}

	// Find mouse/cat/foodPos
	var mi0, ci0, foodPos int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'M' {
				mi0 = idx(i, j)
			}
			if grid[i][j] == 'C' {
				ci0 = idx(i, j)
			}
			if grid[i][j] == 'F' {
				foodPos = idx(i, j)
			}
		}
	}

	// Helper function for finding reachable locations
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	jumpDist := [2]int{}
	jumpDist[mouse] = mouseJump
	jumpDist[cat] = catJump
	reachable := func(ij, turn int) []int {
		i, j := pos(ij)
		var nextPos []int
		dist := jumpDist[turn]
		for _, d := range dirs {
			ii := i
			jj := j
			for k := 1; k <= dist; k++ {
				ii += d[0]
				jj += d[1]
				if !ok(ii, jj) || grid[ii][jj] == '#' {
					break
				}
				nextPos = append(nextPos, idx(ii, jj))
			}
		}
		if turn == cat {
			nextPos = append(nextPos, ij)
		}
		return nextPos
	}

	// Calculate movement options.
	// options[mi][ci][who] = number of available options for 'who' that do not
	// lead to terminal states (wall, catWin or mouseWin).
	options := make([][][2]int, m*n)
	for mi := range options {
		options[mi] = make([][2]int, m*n)
		for ci := range options[mi] {
			options[mi][ci][mouse] = len(reachable(mi, mouse))
			options[mi][ci][cat] = len(reachable(ci, cat))
		}
	}

	// Curr/next holds terminal states, i.e. a game state for which the outcome
	// is known. The idiom of using curr/next and swapping is a memory
	// optimization technique.
	q := [][3]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '#' {
				continue
			}

			ij := idx(i, j)

			// Whenever the cat is on top of the mouse, the game has reached the
			// end.
			q = append(q, [3]int{ij, ij, cat})
			res[ij][ij][mouse] = catWillWin

			if ij == foodPos {
				// Ignore when cat + mouse + food is in the same location
				continue
			}

			// When mouse is on top of food and it's the cat's turn, the game is
			// over, and vice versa
			q = append(q, [3]int{foodPos, ij, cat})
			res[foodPos][ij][cat] = mouseWillWin
			q = append(q, [3]int{foodPos, ij, mouse})
			res[ij][foodPos][mouse] = catWillWin
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		mi := x[0]
		ci := x[1]
		whoseTurn := x[2]

		// If there's a time when it's the cat's turn and the mouse has won,
		// then any prior position when it's the mouse's turn is now a winning
		// position, and vice versa.
		//
		// If there's a time when it's the cat's turn and the cat has won, then
		// any prior position when it's the mouse's turn should avoid this
		// route. If there are no more options available for the mouse for that
		// prior position, then the mouse will lose no matter what, which means
		// that prior position is a losing state.
		//
		result := res[mi][ci][whoseTurn]

		if whoseTurn == cat {
			for _, x := range reachable(mi, mouse) {
				if res[x][ci][mouse] != uncertain {
					// Outcome already known. This can only happen if the mouse
					// has already found an option that leads to a win.
					continue
				}

				// Outcome is not recorded yet
				if result == mouseWillWin {
					// This move led to an undisputable win for the mouse, so
					// the mouse should take it.
					res[x][ci][mouse] = mouseWillWin
					q = append(q, [3]int{x, ci, mouse})
				}
				if result == catWillWin {
					// This move led to a cat win. The mouse should be
					// discouraged to make this move, but there may still be
					// options left.
					options[x][ci][mouse]--
					if options[x][ci][mouse] == 0 {
						// Out of options.. :(
						res[x][ci][mouse] = catWillWin
						q = append(q, [3]int{x, ci, mouse})
					}
				}
			}
		}

		if whoseTurn == mouse {
			for _, y := range reachable(ci, cat) {
				if res[mi][y][cat] != uncertain {
					// Outcome already known. This can only happen if the cat
					// has already found an option that leads to a win.
					continue
				}

				// Outcome is not recorded yet
				if result == catWillWin {
					// This move led to an undisputable win for the cat, so
					// the cat should take it.
					res[mi][y][cat] = catWillWin
					q = append(q, [3]int{mi, y, cat})
				}
				if result == mouseWillWin {
					// This move led to a mouse win. The cat should be
					// discouraged to make this move, but there may still be
					// options left.
					options[mi][y][cat]--
					if options[mi][y][cat] == 0 {
						// Out of options.. :(
						res[mi][y][cat] = mouseWillWin
						q = append(q, [3]int{mi, y, cat})
					}
				}
			}
		}
	}

	return res[mi0][ci0][mouse] == mouseWillWin
}
