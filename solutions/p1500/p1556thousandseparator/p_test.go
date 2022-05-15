package p1556thousandseparator

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	res := make([]byte, 0, 12)
	for i := 0; n > 0; i++ {
		if i > 0 && i%3 == 0 {
			res = append(res, '.')
		}
		res = append(res, byte(n%10+'0'))
		n /= 10
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}
