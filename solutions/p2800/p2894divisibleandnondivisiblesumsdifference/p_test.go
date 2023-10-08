package p2894divisibleandnondivisiblesumsdifference

func differenceOfSums(n int, m int) int {
	var sum1, sum2 int
	for x := 1; x <= n; x++ {
		if x%m == 0 {
			sum2 += x
		} else {
			sum1 += x
		}
	}
	return sum1 - sum2
}
