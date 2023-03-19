package p2595numberofevenandoddbits

import "fmt"

func evenOddBit(n int) []int {
	s := fmt.Sprintf("%b", n)
	res := []int{0, 0}
	for i := range s {
		res[(len(s)-1-i)&1] += int(s[i] - '0')
	}
	return res
}
