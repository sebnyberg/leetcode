package p1281subtractthebroductandsumofdigitsofaninteger

import "fmt"

func subtractProductAndSum(n int) int {
	var sum int
	product := 1
	for _, ch := range fmt.Sprint(n) {
		sum += int(ch - '0')
		product *= int(ch - '0')
	}
	return product - sum
}
