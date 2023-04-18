package p2644findthemaximumdivisibilityscore

func maxDivScore(nums []int, divisors []int) int {
	// There aren't that many, just calculate it
	maxCount := -1
	var res int
	for _, d := range divisors {
		var count int
		for _, x := range nums {
			if x%d == 0 {
				count++
			}
		}
		if count > maxCount || (count == maxCount && res > d) {
			maxCount = count
			res = d
		}
	}
	return res
}
