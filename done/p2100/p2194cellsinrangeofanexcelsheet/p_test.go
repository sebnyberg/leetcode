package p2194cellsinrangeofanexcelsheet

func cellsInRange(s string) []string {
	c1 := s[0]
	r1 := s[1]
	c2 := s[3]
	r2 := s[4]
	var res []string
	for ; c1 <= c2; c1++ {
		for r := r1; r <= r2; r++ {
			res = append(res, string([]byte{c1, r}))
		}
	}
	return res
}
