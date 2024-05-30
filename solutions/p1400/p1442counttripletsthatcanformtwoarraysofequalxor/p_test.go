package p1442counttripletsthatcanformtwoarraysofequalxor

func countTriplets(arr []int) int {
	count := 0
	for i, x := range arr {
		for j, y := range arr[i+1:] {
			x ^= y
			if x == 0 {
				count += j + 1
			}
		}
	}
	return count
}
