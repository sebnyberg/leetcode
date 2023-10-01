package p2765longestalternatingsubarray

func alternatingSubarray(nums []int) int {
	res := -1
	for i := 0; i < len(nums)-1; i++ {
		var want int
		want = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[j-1] != want {
				goto finished
			}
			want = -want
			res = max(res, j-i+1)
		}
	finished:
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
