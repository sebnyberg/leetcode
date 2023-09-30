package p2864maximumoddbinarynumber

import "sort"

func maximumOddBinaryNumber(s string) string {
	// We have to keep one 1 for the end.
	//
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] > bs[j]
	})
	for i := len(s) - 1; i >= 0; i-- {
		if bs[i] == '1' {
			bs[i], bs[len(s)-1] = bs[len(s)-1], bs[i]
			break

		}
	}
	return string(bs)
}
