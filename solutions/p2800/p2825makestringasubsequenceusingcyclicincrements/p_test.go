package p2825makestringasubsequenceusingcyclicincrements

func canMakeSubsequence(str1 string, str2 string) bool {
	// This can be done in a greedy fashion.
	var j int
	for i := range str1 {
		ch := str1[i]
		if ch == str2[j] || byte('a'+(byte(ch-'a')+1)%26) == str2[j] {
			j++
			if j == len(str2) {
				return true
			}
		}
	}
	return false
}
