package p2657findtheprefixcommonarrayoftwoarrays

func findThePrefixCommonArray(A []int, B []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	var inCommon int
	var res []int
	for i := range A {
		if m2[A[i]] > m1[A[i]] {
			inCommon++
		}
		m1[A[i]]++
		if m1[B[i]] > m2[B[i]] {
			inCommon++
		}
		m2[B[i]]++
		res = append(res, inCommon)
	}
	return res
}
