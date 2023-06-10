package p1424diagonaltraverseii

func findDiagonalOrder(nums [][]int) []int {
	var parts [][]int
	for i := range nums {
		for j := range nums[i] {
			d := i + j
			for len(parts) <= d {
				parts = append(parts, []int{})
			}
			parts[d] = append(parts[d], nums[i][j])
		}
	}
	var res []int
	for i := range parts {
		for j := len(parts[i]) - 1; j >= 0; j-- {
			res = append(res, parts[i][j])
		}
	}
	return res
}
