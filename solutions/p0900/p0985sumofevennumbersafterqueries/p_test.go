package p0985sumofevennumbersafterqueries

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	var sumEven int
	for _, x := range nums {
		if x&1 == 0 {
			sumEven += x
		}
	}
	res := make([]int, len(queries))
	for j, q := range queries {
		val, i := q[0], q[1]
		if nums[i]&1 == 0 {
			sumEven -= nums[i]
		}
		nums[i] += val
		if nums[i]&1 == 0 {
			sumEven += nums[i]
		}
		res[j] = sumEven
	}
	return res
}
