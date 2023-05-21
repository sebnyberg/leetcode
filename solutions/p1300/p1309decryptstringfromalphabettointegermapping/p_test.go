package p1309decryptstringfromalphabettointegermapping

func freqAlphabets(s string) string {
	var res []byte
	var i int
	for i < len(s) {
		if i+2 < len(s) && s[i+2] == '#' {
			res = append(res, byte('a'-1+int(s[i]-'0')*10+int(s[i+1]-'0')))
			i += 3
		} else {
			res = append(res, byte('a'-1+int(s[i]-'0')))
			i++
		}
	}
	return string(res)
}
