package p2283checkifnumberhasequaldigitcountanddigitvalue

func digitCount(num string) bool {
	var count [10]int
	for _, x := range num {
		count[x-'0']++
	}
	for i := range num {
		if count[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
