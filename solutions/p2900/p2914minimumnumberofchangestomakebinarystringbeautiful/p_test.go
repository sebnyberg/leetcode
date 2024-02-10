package p2914minimumnumberofchangestomakebinarystringbeautiful

func minChanges(s string) int {
	// It does not matter how many substrings we create, so we may as well
	// pick 0s or 1s for any pair.
	var res int
	for i := 0; i < len(s); i += 2 {
		res += abs(int(s[i+1]-'0') - int(s[i]-'0'))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
