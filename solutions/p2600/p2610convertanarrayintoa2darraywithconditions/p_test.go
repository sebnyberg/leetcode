package p2610convertanarrayintoa2darraywithconditions

func findMatrix(nums []int) [][]int {
	numRow := make(map[int]int)
	var res [][]int
	for _, x := range nums {
		i := numRow[x]
		if i+1 > len(res) {
			res = append(res, []int{})
		}
		res[i] = append(res[i], x)
		numRow[x]++
	}
	return res
}
