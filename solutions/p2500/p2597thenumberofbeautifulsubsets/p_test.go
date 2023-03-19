package p2597thenumberofbeautifulsubsets

import "sort"

func beautifulSubsets(nums []int, k int) int {
	// Just try? There are only up to 20 numbers, 1<<20 is reasonable.
	sort.Ints(nums)
	var res int
	for i := range nums {
		res += dfs(nums, i+1, (1 << i), k)
	}
	return res
}

func dfs(nums []int, i, bm, k int) int {
	if i == len(nums) {
		return 1
	}
	without := dfs(nums, i+1, bm, k)
	x := nums[i]
	for j := 0; j < len(nums); j++ {
		if bm&(1<<j) > 0 && abs(x-nums[j]) == k {
			return without
		}
	}
	with := dfs(nums, i+1, bm|(1<<i), k)
	return without + with
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
