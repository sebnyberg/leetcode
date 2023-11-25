package p1685sumofabsolutedifferencesinasortedarray

func getSumAbsoluteDifferences(nums []int) []int {
	// Since
	//
	// nums.length <= 10^5
	// nums[i] <= 10^4
	//
	// The sum of all elements fits in an integer
	//
	// Hence, we can collect the prefix sum and combine with suffix sum

	n := len(nums)
	pre := make([]int, n+1)
	for i := range nums {
		pre[i+1] = pre[i] + nums[i]
	}
	sum := pre[n]
	suf := sum
	res := make([]int, n)
	for i := range nums {
		suf -= nums[i]
		after := (n - i - 1) * nums[i]
		before := i * nums[i]
		res[i] = suf - after + before - pre[i]
	}
	return res
}
