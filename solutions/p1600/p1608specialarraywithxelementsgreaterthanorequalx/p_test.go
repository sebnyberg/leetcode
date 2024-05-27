package p1608specialarraywithxelementsgreaterthanorequalx

import "sort"

func specialArray(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, x := range nums {
		if x >= i+1 {
			if i == len(nums)-1 || nums[i+1] < i+1 {
				return i + 1
			}
		}
	}
	return -1
}
