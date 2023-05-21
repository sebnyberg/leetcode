package p1346checkifnanditsdoubleexist

func checkIfExist(arr []int) bool {
	m := map[int]int{}
	for _, x := range arr {
		m[x]++
	}
	for x := range m {
		if x == 0 && m[0] >= 2 {
			return true
		}
		if x != 0 && m[x*2] > 0 {
			return true
		}
	}
	return false
}
