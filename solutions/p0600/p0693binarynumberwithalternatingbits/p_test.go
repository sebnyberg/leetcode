package p0693binarynumberwithalternatingbits

func hasAlternatingBits(n int) bool {
	want := n & 1
	for n > 0 {
		if n&1 != want {
			return false
		}
		want = 1 - n&1
		n >>= 1
	}
	return true
}
