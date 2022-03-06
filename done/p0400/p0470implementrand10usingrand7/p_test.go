package p0470implementrand10usingrand7

var rand7 func() int

func rand10() int {
	// Calling rand with sum 10 times yields a number between 1 and 70
	// Divide by 7 to get a result
	res := 40
	for res >= 40 {
		res = 7*(rand7()-1) + rand7() - 1
	}
	return res%10 + 1
}
