package p2661firstcompletelypaintedroworcolumn

func firstCompleteIndex(arr []int, mat [][]int) int {
	count := [2]map[int]int{
		make(map[int]int),
		make(map[int]int),
	}
	m := len(mat)
	n := len(mat[0])
	pos := make(map[int][2]int)
	for i := range mat {
		for j := range mat[i] {
			pos[mat[i][j]] = [2]int{i, j}
		}
	}
	for i, x := range arr {
		p := pos[x]
		count[0][p[0]]++
		count[1][p[1]]++
		if count[0][p[0]] == n || count[1][p[1]] == m {
			return i
		}
	}
	return -1
}
