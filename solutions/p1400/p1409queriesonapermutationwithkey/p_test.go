package p1409queriesonapermutationwithkey

func processQueries(queries []int, n int) []int {
	// Just simulate..
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	m := len(queries)
	res := make([]int, m)
	for i, q := range queries {
		q-- // adjust to zero-index
		var j int
		for nums[j] != q {
			j++
		}
		res[i] = j
		copy(nums[1:], nums[:j])
		nums[0] = q
	}
	return res
}
