package p1250checkifitisagoodarray

func isGoodArray(nums []int) bool {
	// Apparently this is related to Bezout's identity, which states that for
	// any two non-zero integres a and b, there exists integers x and y such
	// that ax + by = gcd(a,b)
	// This tells us that if any subset of nums has a gcd of 1, a solution
	// exists.
	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x = gcd(x, nums[i])
	}
	return x == 1
}
