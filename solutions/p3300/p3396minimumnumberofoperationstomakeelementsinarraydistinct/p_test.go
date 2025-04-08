package p3396minimumnumberofoperationstomakeelementsinarraydistinct

func minimumOperations(nums []int) int {
	m := make(map[int]int)
	var count int
	for i := range nums {
		m[nums[i]]++
		if m[nums[i]] == 2 {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	var k int
	for len(nums) > 0 {
		for k := 0; k < 3 && len(nums) > 0; k++ {
			m[nums[0]]--
			if m[nums[0]] == 1 {
				count--
			}
			nums = nums[1:]
		}
		k++
		if count == 0 {
			return k
		}
	}
	return -1
}
