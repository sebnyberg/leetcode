package p2485findthepivotinteger

func pivotInteger(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	var sum2 int
	for x := 1; x <= n; x++ {
		sum2 += x
		if sum2 == sum-sum2+x {
			return x
		}
	}
	return -1
}
