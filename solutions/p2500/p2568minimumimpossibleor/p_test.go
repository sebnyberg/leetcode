package p2568minimumimpossibleor

func minImpossibleOR(nums []int) int {
	// the answer is simply the first bit not covered by nums.
	var b int
	m := make(map[int]bool)
	for _, x := range nums {
		b |= x
		m[x] = true
	}
	for i := 0; ; i++ {
		if b&(1<<i) == 0 || !m[1<<i] {
			return (1 << i)
		}
	}
}
