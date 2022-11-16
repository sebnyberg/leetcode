package p0982tripletswithbitwiseandequaltozero

func countTriplets(nums []int) int {
	// Let's do a dumb implementation first to see.
	ones := make(map[int]int)
	twos := make(map[int]int)
	var res int
	for _, x := range nums {
		for y, v := range twos {
			if y&x == 0 {
				res += 3 * v
			}
		}
		for y, v := range ones {
			twos[y&x] += 2 * v
		}
		twos[x&x] += 2
		ones[x]++
		if x&x&x == 0 {
			res++
		}
	}
	return res
}
