package p0287missingnumber

var _ = missingNumber

func missingNumber(nums []int) (res int) {
	for i, n := range nums {
		res ^= n
		res ^= i + 1
	}
	return
}
