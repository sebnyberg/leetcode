package bitmask

func debugBitMask(bm uint32, n uint8) string {
	b := uint32(2)
	out := make([]byte, 0, n)
	for i := uint8(0); i < n; i++ {
		if bm&b > 0 {
			out = append(out, 'X')
		} else {
			out = append(out, '-')
		}
		b <<= 1
	}
	return string(out)
}
