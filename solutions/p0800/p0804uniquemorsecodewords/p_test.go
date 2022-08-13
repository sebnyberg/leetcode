package p0804uniquemorsecodewords

var mapping = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]struct{})
	var buf []byte
	morseify := func(s string) string {
		buf = buf[:0]
		for _, ch := range s {
			buf = append(buf, mapping[ch-'a']...)
		}
		return string(buf)
	}
	for _, w := range words {
		m[morseify(w)] = struct{}{}
	}
	return len(m)
}
