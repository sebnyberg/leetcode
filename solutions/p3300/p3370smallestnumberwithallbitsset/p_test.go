package p3370smallestnumberwithallbitsset

func smallestNumber(n int) int {
	for n&(n-1) > 0 {
		n &= n - 1
	}
	return (n << 1) - 1
}
