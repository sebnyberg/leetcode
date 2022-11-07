package p1323maximum69number

func maximum69Number(num int) int {
	p := func(s string) int {
		x, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		return int(x)
	}
	s := fmt.Sprint(num)
	maxVal := p(s)
	for i := range s {
		mid := "9"
		if s[i] == '9' {
			mid = "6"
		}
		d := p(s[:i] + mid + s[i+1:])
		if d > maxVal {
			maxVal = d
		}
	}
	return maxVal
}
