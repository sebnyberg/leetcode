package p2427numberofcommonfactors

func commonFactors(a int, b int) int {
	var res int
	for x := 1; x <= a; x++ {
		if a%x == 0 && b%x == 0 {
			res++
		}
	}
	return res
}
