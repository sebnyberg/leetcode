package p3000maximumareaoflongestdiagonalrectangle

func areaOfMaxDiagonal(dimensions [][]int) int {
	var maxDiag int
	var res int
	for _, d := range dimensions {
		diag := d[0]*d[0] + d[1]*d[1]
		area := d[0] * d[1]
		if diag > maxDiag || (diag == maxDiag && area > res) {
			maxDiag = diag
			res = area
		}
	}
	return res
}
