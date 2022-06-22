package p0754reachanumber

func reachNumber(target int) int {
	// 1 left 2 right => 1
	// 1 right 2 right => 3
	// 1 right 2 left => -1
	// {1, 3, -1, -3}
	// {4, 6, 2, 0}, {-2, 0, -4, -6}
	// sum = n * (n+1) / 2
	sum := func(x int) int {
		return x * (x + 1) / 2
	}
	if target < 0 {
		target = -target
	}
	lo, hi := 0, target+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		s := sum(mid)
		if s >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	x := lo
	for sum(x)&1 != target&1 {
		x++
	}
	return x
}
