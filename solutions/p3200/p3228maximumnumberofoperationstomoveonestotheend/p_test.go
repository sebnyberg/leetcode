package p3228maximumnumberofoperationstomoveonestotheend

func maxOperations(s string) int {
	var count int
	var res int
	for i := 0; i <= len(s); i++ {
		if i > 0 && (i == len(s) || s[i] == '1') && s[i-1] == '0' {
			res += count
		}
		if i < len(s) && s[i] == '1' {
			count++
		}
	}
	return res
}
