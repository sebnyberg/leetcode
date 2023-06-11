package p2734lexicographicallysmalleststringaftersubstringoperation

func smallestString(s string) string {
	// Include sharacters until a character is 'a'
	res := []byte(s)
	var l int
	var changes int
	for l < len(res) && res[l] == 'a' {
		l++
	}
	for l < len(res) && res[l] != 'a' {
		res[l]--
		changes++
		l++
	}
	if changes == 0 {
		res[len(res)-1] = 'z'
	}
	return string(res)
}
