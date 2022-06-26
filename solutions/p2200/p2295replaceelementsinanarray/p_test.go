package p2295replaceelementsinanarray

func arrayChange(nums []int, operations [][]int) []int {
	n := len(nums)
	m := make(map[int]int, n)
	cpy := make([]int, n)
	copy(cpy, nums)
	nums = cpy
	for i, num := range nums {
		m[num] = i
	}
	for _, op := range operations {
		a, b := op[0], op[1]
		if _, exists := m[a]; !exists {
			continue
		}
		nums[m[a]] = b
		m[b] = m[a]
		delete(m, a)
	}
	return nums
}
