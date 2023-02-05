package p2553separatethedigitsinanarray

import "fmt"

func separateDigits(nums []int) []int {
	var res []int

	for _, x := range nums {
		for _, ch := range fmt.Sprint(x) {
			res = append(res, int(ch-'0'))
		}
	}
	return res
}
