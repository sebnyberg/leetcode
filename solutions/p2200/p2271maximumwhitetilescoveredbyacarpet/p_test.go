package p2271maximumwhitetilescoveredbyacarpet

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumWhiteTiles(t *testing.T) {
	for _, tc := range []struct {
		tiles     [][]int
		carpetLen int
		want      int
	}{
		{
			leetcode.ParseMatrix("[[3745,3757],[3663,3681],[3593,3605],[3890,3903],[3529,3539],[3684,3686],[3023,3026],[2551,2569],[3776,3789],[3243,3256],[3477,3497],[2650,2654],[2264,2266],[2582,2599],[2846,2863],[2346,2364],[3839,3842],[3926,3935],[2995,3012],[3152,3167],[4133,4134],[4048,4058],[3719,3730],[2498,2510],[2277,2295],[4117,4128],[3043,3054],[3394,3402],[3921,3924],[3500,3514],[2789,2808],[3291,3294],[2873,2881],[2760,2760],[3349,3362],[2888,2899],[3802,3822],[3540,3542],[3128,3142],[2617,2632],[3979,3994],[2780,2781],[3213,3233],[3099,3113],[3646,3651],[3956,3963],[2674,2691],[3860,3873],[3363,3370],[2727,2737],[2453,2471],[4011,4031],[3566,3577],[2705,2707],[3560,3565],[3454,3456],[3655,3660],[4100,4103],[2382,2382],[4032,4033],[2518,2531],[2739,2749],[3067,3079],[4068,4074],[2297,2312],[2489,2490],[2954,2974],[2400,2418],[3271,3272],[3628,3632],[3372,3377],[2920,2940],[3315,3330],[3417,3435],[4146,4156],[2324,2340],[2426,2435],[2373,2376],[3621,3626],[2826,2832],[3937,3949],[3178,3195],[4081,4082],[4092,4098],[3688,3698]]"),
			1638,
			822,
		},
		{
			leetcode.ParseMatrix("[[8051,8057],[8074,8089],[7994,7995],[7969,7987],[8013,8020],[8123,8139],[7930,7950],[8096,8104],[7917,7925],[8027,8035],[8003,8011]]"),
			9854,
			126,
		},
		{
			leetcode.ParseMatrix("[[1,5],[10,11],[12,18],[20,25],[30,32]]"),
			10,
			9,
		},
		{
			leetcode.ParseMatrix("[[10,11],[1,1]]"),
			2,
			2,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tiles), func(t *testing.T) {
			require.Equal(t, tc.want, maximumWhiteTiles(tc.tiles, tc.carpetLen))
		})
	}
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	// Conjecture: a carpet must (should) either start at the start of a tile, or
	// end at the end of a tile.

	// Sort just in case
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	for i := range tiles {
		tiles[i][1] += 1
	}
	// n := len(tiles)
	// tiles = append(tiles, []int{tiles[n-1][1] * 2, tiles[n-1][1]*2 + 1})

	q := [][]int{}
	var dist int
	var res int
	for i, t := range tiles {
		_ = i
		for len(q) > 0 && t[0]-q[0][1] >= carpetLen {
			dist -= (q[0][1] - q[0][0])
			q = q[1:]
		}
		q = append(q, t)

		// Place mat at the start of the first tile
		remains := max(0, carpetLen-(t[0]-q[0][0]))
		b := min(remains, t[1]-t[0])
		res = max(res, dist+b)

		// blabla
		for len(q) > 0 && t[1]-q[0][1] >= carpetLen {
			dist -= (q[0][1] - q[0][0])
			q = q[1:]
		}
		remains = max(0, carpetLen-(t[0]-q[0][0]))
		b = min(remains, t[1]-t[0])
		res = max(res, dist+b)

		dist += t[1] - t[0]
		for len(q) > 0 && q[len(q)-1][1]-q[0][0] >= carpetLen {
			dist -= (q[0][1] - q[0][0])
			q = q[1:]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
