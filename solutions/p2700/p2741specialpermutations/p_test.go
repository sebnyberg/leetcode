package p2741specialpermutations

const mod = 1e9 + 7

func specialPerm(nums []int) int {
	// A zero modulo means that the two numbers share a prime factor
	// The question is how we explore all possible permutations of numbers.
	// Since there are only 14 numbers in total, we could probably do exhaustive
	// DFS with memoization
	mem := make(map[[2]int]int)
	var res int
	for i := 0; i < len(nums); i++ {
		res = (res + dfs(mem, nums, nums[i], (1<<i))) % mod
	}
	return res
}

func dfs(mem map[[2]int]int, nums []int, prev, bm int) int {
	n := len(nums)
	if bm == (1<<n)-1 {
		// All numbers are covered
		return 1
	}
	key := [2]int{prev, bm}
	if v, exists := mem[key]; exists {
		return v
	}
	var res int
	for i := 0; i < n; i++ {
		if bm&(1<<i) > 0 || (nums[i]%prev != 0 && prev%nums[i] != 0) {
			continue
		}
		res += dfs(mem, nums, nums[i], bm|(1<<i))
	}
	mem[key] = res % mod
	return res
}
