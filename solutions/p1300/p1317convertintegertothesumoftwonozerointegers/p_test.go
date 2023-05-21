package p1317convertintegertothesumoftwonozerointegers

func getNoZeroIntegers(n int) []int {
	// Let's do something lazy: do half then move from one to the other until
	// there are no zeroes.
	left := n / 2
	right := n/2 + n&1
	haszeroes := func(x int) bool {
		for x > 0 {
			if x%10 == 0 {
				return true
			}
			x /= 10
		}
		return false
	}
	for haszeroes(left) || haszeroes(right) {
		left--
		right++
	}
	return []int{left, right}
}
