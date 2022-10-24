package p2447numberofsubarrayswithgcdequaltok

func subarrayGCD(nums []int, k int) int {
	curr := make(map[int]int)
	next := make(map[int]int)
	var res int
	for _, x := range nums {
		for k := range next {
			delete(next, k)
		}
		for val, count := range curr {
			next[gcd(val, x)] += count
		}
		next[x]++
		res += next[k]
		curr, next = next, curr
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
