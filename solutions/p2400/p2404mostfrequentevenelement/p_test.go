package p2404mostfrequentevenelement

func mostFrequentEven(nums []int) int {
	m := make(map[int]int)
	for _, x := range nums {
		if x&1 == 0 {
			m[x]++
		}
	}
	var maxCount int
	var maxVal int
	for val, count := range m {
		if count > maxCount || (count == maxCount && val < maxVal) {
			maxCount = count
			maxVal = val
		}
	}
	if len(m) == 0 {
		return -1
	}
	return maxVal
}
