package p1033movingstonesuntilconsecutive

import "sort"

func numMovesStones(a int, b int, c int) []int {
	nums := []int{a, b, c}
	sort.Ints(nums)
	minMoves := 2
	if nums[0] == nums[1]-1 && nums[1] == nums[2]-1 {
		minMoves = 0
	} else if nums[1]-nums[0] <= 2 || nums[2]-nums[1] <= 2 {
		minMoves = 1
	}
	maxMoves := nums[1] - nums[0] - 1 + nums[2] - nums[1] - 1
	return []int{minMoves, maxMoves}
}
