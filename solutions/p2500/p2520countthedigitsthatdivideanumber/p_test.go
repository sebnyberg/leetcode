package p2520countthedigitsthatdivideanumber

import "fmt"

func countDigits(num int) int {
	var res int
	for _, ch := range fmt.Sprint(num) {
		if num%int(ch-'0') == 0 {
			res++
		}
	}
	return res
}
