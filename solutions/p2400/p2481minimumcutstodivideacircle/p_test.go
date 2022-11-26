package p2481minimumcutstodivideacircle

func numberOfCuts(n int) int {
	// A full cut increases the number of slices by twice
	// A half-cut increases by 1
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}
