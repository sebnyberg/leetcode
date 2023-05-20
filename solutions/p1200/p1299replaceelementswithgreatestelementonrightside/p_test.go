package p1299replaceelementswithgreatestelementonrightside

func replaceElements(arr []int) []int {
	greatest := -1
	for i := len(arr) - 1; i >= 0; i-- {
		greatest, arr[i] = max(greatest, arr[i]), greatest
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
