package p1512numberofgoodpairs

func numIdenticalPairs(nums []int) int {
	count := make(map[int]int)
	var res int
	for _, x := range nums {
		res += count[x]
		count[x]++
	}
	return res
}
