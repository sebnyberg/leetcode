package p1414findtheminimumnumberoffibonaccinumberswhosesumisk

func findMinFibonacciNumbers(k int) int {
	fib := []int{1, 1}
	var j int
	for fib[j]+fib[j+1] <= k {
		fib = append(fib, fib[j]+fib[j+1])
		j++
	}
	var res int
	for r := len(fib) - 1; k > 0; r-- {
		if k >= fib[r] {
			res++
			k -= fib[r]
		}
	}
	return res
}
