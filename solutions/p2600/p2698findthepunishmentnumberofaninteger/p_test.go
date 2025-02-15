package p2698findthepunishmentnumberofaninteger

import "fmt"

func punishmentNumber(n int) int {
	var res int
	for x := 1; x <= n; x++ {
		square := fmt.Sprint(x * x)
		if findSum(square, 0, 0, x) {
			res += x * x
		}
	}
	return res
}

func findSum(square string, buf, i, val int) bool {
	if val < 0 {
		return false
	}
	buf += int(square[i] - '0')
	if i == len(square)-1 {
		return val-buf == 0
	}
	return findSum(square, buf*10, i+1, val) || findSum(square, 0, i+1, val-buf)
}
