package p2799countcompletesubarraysinanarray

func countCompleteSubarrays(nums []int) int {
	m := make(map[int]int)
	n := len(nums)
	for _, x := range nums {
		m[x] = 1
	}
	ndistinct := len(m)
	for k := range m {
		delete(m, k)
	}
	var k int
	var l int
	var res int
	for r, x := range nums {
		m[x]++
		if m[x] == 1 {
			k++
		}
		for k == ndistinct {
			res += n - r
			// move left pointer
			m[nums[l]]--
			if m[nums[l]] == 0 {
				k--
			}
			l++
		}
	}
	return res
}
