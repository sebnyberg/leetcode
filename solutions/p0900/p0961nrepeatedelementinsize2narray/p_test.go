package p0961nrepeatedelementinsize2narray

func repeatedNTimes(nums []int) int {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}
	n := len(nums) / 2
	for x, c := range m {
		if c == n {
			return x
		}
	}
	return 0
}
