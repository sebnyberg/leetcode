package p1027longestarithmeticsubsequence

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	idx := make([]int, 501)
	m := make([]int, n)
	var res int
	for d := -500; d <= 500; d++ {
		for i := range idx {
			idx[i] = -1
		}
		for i := range m {
			m[i] = 1
		}
		for i, x := range nums {
			prev := x - d
			if prev >= 0 && prev <= 500 && idx[prev] != -1 {
				m[i] = m[idx[prev]] + 1
			}
			idx[x] = i
			res = max(res, m[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
