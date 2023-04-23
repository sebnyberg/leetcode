package p2652summultiples

func sumOfMultiples(n int) int {
	var res int
	for x := 1; x <= n; x++ {
		if x%3 == 0 || x%5 == 0 || x%7 == 0 {
			res += x
		}
	}
	return res
}
