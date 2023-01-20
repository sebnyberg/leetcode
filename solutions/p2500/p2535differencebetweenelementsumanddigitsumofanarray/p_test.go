package p2535differencebetweenelementsumanddigitsumofanarray

import "fmt"

func differenceOfSum(nums []int) int {
	var sum int
	var dsum int
	for _, x := range nums {
		for _, ch := range fmt.Sprint(x) {
			dsum += int(ch - '0')
		}
		sum += x
	}
	return abs(sum - dsum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
