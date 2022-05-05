package p0412fizzbuzz

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res[i-1] = "FizzBuzz"
		case i%3 == 0:
			res[i-1] = "Fizz"
		case i%5 == 0:
			res[i-1] = "Buzz"
		default:
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}
