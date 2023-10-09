package p2780minimumindexofavalidsplit

func minimumIndex(nums []int) int {
	// An element being dominant means that it holds the majority count.
	//
	// So the goal is to find the first place where the same element holds the
	// majority count on both sides.
	//
	count := make(map[int]int)
	n := len(nums)
	for i := range nums {
		count[nums[i]]++
	}
	want := -1
	for x, c := range count {
		if c*2 > n {
			want = x
			break
		}
	}
	if want == -1 {
		return -1
	}

	leftCount := make(map[int]int)
	for i := range nums {
		count[nums[i]]--
		leftCount[nums[i]]++
		r := n - (i + 1)
		if leftCount[want]*2 > i+1 && count[want]*2 > r {
			return i
		}
	}
	return -1
}
