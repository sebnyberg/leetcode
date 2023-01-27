package p2544alternatingdigitsum

import "fmt"

func alternateDigitSum(n int) int {
	var res int
	sign := 1
	for _, ch := range fmt.Sprint(n) {
		res += sign * int(ch-'0')
		sign *= -1
	}
	return res
}
