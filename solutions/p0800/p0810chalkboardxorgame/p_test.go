package p0810chalkboardxorgame

func xorGame(nums []int) bool {
	var xor int
	for _, x := range nums {
		xor ^= x
	}
	if xor == 0 {
		return true
	}
	return len(nums)&1 == 0
}
