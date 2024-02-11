package p2917findthekorofanarray

func findKOr(nums []int, k int) int {
	var count [32]int
	for _, x := range nums {
		var i int
		for x > 0 {
			count[i] += x & 1
			x >>= 1
			i++
		}
	}
	var res int
	for i := 1; i <= 32; i++ {
		if count[i-1] >= k {
			res |= (1 << (i - 1))
		}
	}
	return res
}
