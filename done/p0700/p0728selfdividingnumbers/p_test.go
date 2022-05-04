package p0728selfdividingnumbers

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	check := func(x int) bool {
		for _, ch := range fmt.Sprint(x) {
			if ch == '0' {
				return false
			}
			if x%int(ch-'0') != 0 {
				return false
			}
		}
		return true
	}
	res := make([]int, 0, right-left+1)
	for x := left; x <= right; x++ {
		if check(x) {
			res = append(res, x)
		}
	}
	return res
}
