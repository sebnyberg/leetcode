package p2643rowwithmaximumones

func rowAndMaximumOnes(mat [][]int) []int {
	res := []int{-1, -1}
	for i := range mat {
		var count int
		for j := range mat[i] {
			count += mat[i][j]
		}
		if count > res[1] {
			res = []int{i, count}
		}
	}
	return res
}
