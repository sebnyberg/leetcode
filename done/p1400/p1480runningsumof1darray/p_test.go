package p1480runningsumof1darray

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	for i, n := range nums {
		res[i] = n
		if i > 0 {
			res[i] += res[i-1]
		}
	}
	return res
}
