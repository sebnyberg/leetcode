package p0873lengthoflongestfibonaccisubsequence

func lenLongestFibSubseq(arr []int) int {
	seen := make(map[int]struct{})
	for _, x := range arr {
		seen[x] = struct{}{}
	}
	ok := func(sum int) bool {
		_, exists := seen[sum]
		return exists
	}
	var res int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a := arr[i]
			b := arr[j]
			n := 2
			for ok(a + b) {
				n++
				a, b = b, a+b
				res = max(res, n)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
