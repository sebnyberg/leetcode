package p2962countsubarrayswheremaxelementappearsatleastktimes

func countSubarrays(nums []int, k int) int64 {
	var maxNum int
	for _, x := range nums {
		maxNum = max(maxNum, x)
	}

	var pos []int
	var res int
	for i, x := range nums {
		if x == maxNum {
			pos = append(pos, i)
			if len(pos) > k {
				pos = pos[1:]
			}
		}
		if len(pos) == k {
			res += pos[0] + 1
		}
	}
	return int64(res)
}
