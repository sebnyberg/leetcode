package p2455averagevalueofevennumbersthataredivisiblebythree

func averageValue(nums []int) int {
	var sum, count int
	for _, x := range nums {
		if x%2 == 0 && x%3 == 0 {
			sum += x
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
