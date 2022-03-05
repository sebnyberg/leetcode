package p0557reversewordsinastring3

import "strings"

func reverseWords(s string) string {
	parts := strings.Fields(s)
	for i := range parts {
		bs := []byte(parts[i])
		for l, r := 0, len(parts[i])-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		parts[i] = string(bs)
	}
	return strings.Join(parts, " ")
}
