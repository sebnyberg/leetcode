package p2899lastvisitedintegers

import "strconv"

func lastVisitedIntegers(words []string) []int {
	var nums []int
	var res []int
	var prevCount int
	for _, w := range words {
		if w != "prev" {
			prevCount = 0
			x, err := strconv.Atoi(w)
			if err != nil {
				panic(err)
			}
			nums = append(nums, x)
			continue
		}
		prevCount++
		if prevCount <= len(nums) {
			n := len(nums)
			res = append(res, nums[n-prevCount])
		} else {
			res = append(res, -1)
		}
	}
	return res
}
