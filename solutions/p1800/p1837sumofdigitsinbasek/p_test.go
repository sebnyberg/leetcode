package p1837sumofdigitsinbasek

func sumBase(n int, k int) int {
	res := 0
	for n > 0 {
		res += n % k
		n /= k
	}
	return res
}
