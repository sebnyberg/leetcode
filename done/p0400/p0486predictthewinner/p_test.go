package p0486predictthewinner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PredictTheWinner(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{0}, true},
		{[]int{1, 1}, true},
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, PredictTheWinner(tc.nums))
		})
	}
}

func PredictTheWinner(nums []int) bool {
	mem := make(map[[4]int]int)
	res := dfs(mem, nums, 0, len(nums)-1, [2]int{}, 0)
	return res == 0
}

func dfs(mem map[[4]int]int, nums []int, l, r int, scores [2]int, player int) (winner int) {
	player &= 1

	// Done
	if r < l {
		other := (player + 1) & 1
		if scores[player] == scores[other] {
			return 0
		} else if scores[player] >= scores[other] {
			return player
		} else {
			return (player + 1) & 1
		}
	}

	// Current player wins if one move exists that grants a win
	k := [4]int{l, r, scores[0], scores[1]}
	if winningPlayer, exists := mem[k]; exists {
		return winningPlayer
	}

	// Try to remove left
	scores[player] += nums[l]
	if dfs(mem, nums, l+1, r, scores, player+1) == player {
		mem[k] = player
		return mem[k]
	}
	scores[player] -= nums[l]
	scores[player] += nums[r]
	if dfs(mem, nums, l, r-1, scores, player+1) == player {
		mem[k] = player
		return mem[k]
	}
	mem[k] = (player + 1) % 2
	return mem[k]
}
