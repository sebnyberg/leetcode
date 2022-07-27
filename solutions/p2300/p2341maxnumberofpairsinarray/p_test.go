package p2341maxnumberofpairsinarray

func numberOfPairs(nums []int) []int {
	var count [101]int
	res := []int{0, 0}
	for _, x := range nums {
		count[x]++
		if count[x] == 2 {
			res[0]++
			count[x] = 0
		}
	}
	for _, c := range count {
		res[1] += c
	}
	return res
}
