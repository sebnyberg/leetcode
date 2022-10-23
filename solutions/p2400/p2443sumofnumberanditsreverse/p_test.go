package p2443sumofnumberanditsreverse

func sumOfNumberAndReverse(num int) bool {
	if num == 0 {
		return true
	}
	rev := func(x int) int {
		var res int
		for x > 0 {
			res *= 10
			res += x % 10
			x /= 10
		}
		return res
	}
	// Try every number from 0 to num/2
	for x := num - 1; x >= num/2; x-- {
		if rev(x)+x == num {
			return true
		}
	}
	return false
}
