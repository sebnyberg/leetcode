package p1523countoddnumbersinanintervalrange

func countOdds(low int, high int) int {
	n := high - low + 1
	return (n / 2) + (n & low & 1)
}
