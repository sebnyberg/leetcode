package p3105longeststrictlyincreasingorstrictlydecreasingsubarray

func longestMonotonicSubarray(nums []int) int {
	res := 1
	m := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			m++
		} else {
			m = 1
		}
		res = max(res, m)
	}
	m = 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			m++
		} else {
			m = 1
		}
		res = max(res, m)
	}
	return res
}
