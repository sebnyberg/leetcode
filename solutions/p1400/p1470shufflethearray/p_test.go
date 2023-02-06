package p1470shufflethearray

func shuffle(nums []int, n int) []int {
	res := make([]int, n*2)
	for i := 0; i < n; i++ {
		res[i*2] = nums[i]
		res[i*2+1] = nums[n+i]
	}
	return res
}
