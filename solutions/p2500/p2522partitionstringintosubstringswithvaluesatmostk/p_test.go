package p2522partitionstringintosubstringswithvaluesatmostk

func minimumPartition(s string, k int) int {
	// Just do a greedy partitioning
	var res int
	var n int
	var val int
	for i := range s {
		v := int(s[i] - '0')
		if v > k {
			return -1
		}
		if val*10+v > k {
			n = 0
			val = 0
			res++
		}
		n++
		val = val*10 + v
	}
	return res + 1
}
