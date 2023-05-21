package p1343numbesofsubarraysofsizekandaveragegreaterthanorequaltothreshold

func numOfSubarrays(arr []int, k int, threshold int) int {
	var sum int
	var res int
	for i, x := range arr {
		sum += x
		if i >= k {
			sum -= arr[i-k]
		}
		if i >= k-1 && float64(sum)/float64(k) >= float64(threshold) {
			res++
		}
	}
	return res
}
