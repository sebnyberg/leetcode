package p2486appendcharacterstostringtomakesubsequence

func appendCharacters(s string, t string) int {
	var j int
	for i := range s {
		if s[i] == t[j] {
			j++
		}
		if j == len(t) {
			break
		}
	}
	return len(t) - j
}
