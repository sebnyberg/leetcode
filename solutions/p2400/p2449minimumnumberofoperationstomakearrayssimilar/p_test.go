package p2449minimumnumberofoperationstomakearrayssimilar

import "sort"

func makeSimilar(nums []int, target []int) int64 {
	n := len(nums)
	evenNums := make([]int, 0, n)
	evenTargets := make([]int, 0, n)
	oddNums := make([]int, 0, n)
	oddTargets := make([]int, 0, n)
	for i := range nums {
		if nums[i]%2 == 0 {
			evenNums = append(evenNums, nums[i])
		} else {
			oddNums = append(oddNums, nums[i])
		}
		if target[i]%2 == 0 {
			evenTargets = append(evenTargets, target[i])
		} else {
			oddTargets = append(oddTargets, target[i])
		}
	}

	sort.Ints(evenNums)
	sort.Ints(oddNums)
	sort.Ints(evenTargets)
	sort.Ints(oddTargets)
	var ops int64
	for i := range evenNums {
		ops += int64(abs(evenNums[i]-evenTargets[i]) / 2)
	}
	for i := range oddNums {
		ops += int64(abs(oddNums[i]-oddTargets[i]) / 2)
	}
	return ops / 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
