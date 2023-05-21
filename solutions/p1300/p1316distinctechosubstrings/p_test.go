package p1316distinctechosubstrings

func distinctEchoSubstrings(text string) int {
	// Note: a substring is a contiguous sequence
	// This is not a hard problem? Just run a sliding window of all sizes <=
	// len(text)/2 then check if the right side is equal to the left..
	m := make(map[string]struct{})
	for k := len(text) / 2; k >= 1; k-- {
		for i := k; i <= len(text)-k; i++ {
			if text[i-k:i] == text[i:i+k] {
				m[text[i:i+k]] = struct{}{}
			}
		}
	}
	return len(m)
}
