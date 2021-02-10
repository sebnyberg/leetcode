package p0741cherrypickup

// func Test_cherryPickup(t *testing.T) {
// 	for _, tc := range []struct {
// 		grid [][]int
// 		want int
// 	}{
// 		{},
// 	} {
// 		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
// 			require.Equal(t, tc.want, cherryPickup(tc.grid))
// 		})
// 	}
// }

// func cherryPickup(grid [][]int) int {
// 	// Plan: create cherries grid
// 	// Pick up the maximum amount of cherries until the end of the grid
// 	m, n := len(grid), len(grid[0])
// 	maxCherries := make([][]int, m+1)
// 	for i := range maxCherries {
// 		maxCherries[i] = make([]int, n+1)
// 	}

// 	pick(grid, maxCherries)
// }

// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			left, above := maxCherries[i][j-1], maxCherries[i-1][j]
// 			if left == above {
// 				switch left {
// 				case -1:
// 					maxCherries[i][j] = -1
// 				case 0:
// 					maxCherries[i][j] = 0
// 				default:
// 					// The choice is ambiguous
// 					// Branch out and try both alternatives

// 				}
// 			}
// 		}
// 	}
