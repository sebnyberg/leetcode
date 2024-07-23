package p1636sortarraybyincreasingfrequency

import "sort"

func frequencySort(nums []int) []int {
	var freq [201]int
	for _, x := range nums {
		freq[x+100]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]+100] == freq[nums[j]+100] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]+100] < freq[nums[j]+100]
	})
	return nums
}
