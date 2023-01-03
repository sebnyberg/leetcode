package p1072flipcolumnsformaximumnumberofequalrows

func maxEqualRowsAfterFlips(matrix [][]int) int {
	var s []byte
	var s2 []byte
	score := make(map[string]int)
	for i := range matrix {
		s = s[:0]
		s2 = s2[:0]
		for _, v := range matrix[i] {
			s = append(s, byte('0'+v))
			s2 = append(s2, byte('0'+(1-v)))
		}
		score[string(s)]++
		score[string(s2)]++
	}
	var res int
	for _, c := range score {
		res = max(res, c)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
