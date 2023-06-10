package p2729checkifthenumberisfascinating

import "fmt"

func isFascinating(n int) bool {
	res := fmt.Sprint(n) + fmt.Sprint(n*2) + fmt.Sprint(n*3)
	var count [10]int
	for i := range res {
		if res[i] == '0' {
			return false
		}
		if count[res[i]-'0'] > 0 {
			return false
		}
		count[res[i]-'0'] = 1
	}
	for i := 1; i <= 9; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return true
}
