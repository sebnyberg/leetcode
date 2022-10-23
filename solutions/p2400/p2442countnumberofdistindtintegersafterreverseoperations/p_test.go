package p2442countnumberofdistindtintegersafterreverseoperations

func countDistinctIntegers(nums []int) int {
	m := make(map[int]struct{})
	rev := func(x int) int {
		var res int
		for x > 0 {
			res *= 10
			res += x % 10
			x /= 10
		}
		return res
	}
	for _, x := range nums {
		m[x] = struct{}{}
		m[rev(x)] = struct{}{}
	}
	return len(m)
}
