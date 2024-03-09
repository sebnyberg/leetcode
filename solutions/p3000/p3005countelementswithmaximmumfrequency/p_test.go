package p3005countelementswithmaximmumfrequency

import "sort"

func maxFrequencyElements(nums []int) int {
	m := map[int]int{}
	type numFreq struct {
		num  int
		freq int
	}
	freqs := []numFreq{}
	for i := range nums {
		if _, exists := m[nums[i]]; !exists {
			m[nums[i]] = len(freqs)
			freqs = append(freqs, numFreq{num: nums[i], freq: 1})
		} else {
			freqs[m[nums[i]]].freq++
		}
	}
	maxFreq := freqs[0].freq
	sort.Slice(freqs, func(i, j int) bool {
		maxFreq = max(maxFreq, max(freqs[i].freq, freqs[j].freq))
		return freqs[i].freq > freqs[j].freq
	})
	var res int
	for i := range freqs {
		if freqs[i].freq < maxFreq {
			break
		}
		res += maxFreq
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
