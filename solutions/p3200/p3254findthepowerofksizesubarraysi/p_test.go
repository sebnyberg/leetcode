package p3254findthepowerofksizesubarraysi

func resultsArray(nums []int, k int) []int {
	var l int
	var res []int
	for i, x := range nums {
		if i == 0 || nums[i] != nums[i-1]+1 {
			l = i
		}
		if i < k-1 {
			continue
		}
		if i-l+1 >= k {
			res = append(res, x)
		} else {
			res = append(res, -1)
		}
	}
	return res
}
