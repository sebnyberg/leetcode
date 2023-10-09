package p2778sumofsquaresofspecialelements

func sumOfSquares(nums []int) int {
	n := len(nums)
	var res int
	for i, x := range nums {
		if n%(i+1) == 0 {
			res += x * x
		}
	}
	return res
}
