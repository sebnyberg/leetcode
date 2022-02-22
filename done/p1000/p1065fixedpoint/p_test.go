package p1065fixedpoint

func fixedPoint(arr []int) int {
	for i, v := range arr {
		if v == i {
			return i
		}
	}
	return -1
}
