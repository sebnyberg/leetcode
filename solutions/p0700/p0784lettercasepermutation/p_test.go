package p0784lettercasepermutation

import "unicode"

func letterCasePermutation(s string) []string {
	res := make([][]byte, 0)
	res = append(res, []byte{})
	for j, ch := range s {
		if unicode.IsDigit(ch) {
			for i := range res {
				res[i] = append(res[i], byte(ch))
			}
		} else {
			n := len(res)
			ch := unicode.ToLower(ch)
			for i := 0; i < n; i++ {
				res[i] = append(res[i], byte(ch))
				res = append(res, make([]byte, j))
				copy(res[n+i], res[i])
				res[n+i] = append(res[n+i], byte(unicode.ToUpper(ch)))
			}
		}
	}
	ss := make([]string, len(res))
	for i := range ss {
		ss[i] = string(res[i])
	}
	return ss
}
