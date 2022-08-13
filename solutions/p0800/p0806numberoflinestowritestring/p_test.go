package p0806numberoflinestowritestring

func numberOfLines(widths []int, s string) []int {
	var w [256]byte
	var ww byte
	l := 1
	for i, width := range widths {
		w['a'+i] = byte(width)
	}
	for _, ch := range s {
		if ww+w[ch] > 100 {
			l++
			ww = 0
		}
		ww += w[ch]
	}
	return []int{l, int(ww)}
}
