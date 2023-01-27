package p2546applybitwiseoperationstomakestringsequal

func makeStringsEqual(s string, target string) bool {
	// Any 0 1 can be changed do 1 1
	// Any 1 1 can be changed to 0 1
	// The last 1 cannot be removed
	var got [2]int
	var want [2]int
	for i := range s {
		got[s[i]-'0']++
		want[target[i]-'0']++
	}
	return (want[1] > 0 && got[1] > 0) ||
		want[1] == 0 && got[1] == 0
}
