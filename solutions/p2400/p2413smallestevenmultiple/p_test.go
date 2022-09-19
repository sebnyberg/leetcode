package p2413smallestevenmultiple

func smallestEvenMultiple(n int) int {
	if n&1 == 1 {
		return n * 2
	}
	return n
}
