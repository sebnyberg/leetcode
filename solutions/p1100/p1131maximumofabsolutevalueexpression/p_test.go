package p1131maximumofabsolutevalueexpression

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	var res int
	for i := range arr1 {
		for j := i + 1; j < len(arr1); j++ {
			a := abs(arr1[j] - arr1[i])
			b := abs(arr2[j] - arr2[i])
			res = max(res, a+b+(j-i))
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
