package p0860lemonadechange

func lemonadeChange(bills []int) bool {
	const five = 0
	const ten = 1
	var count [2]int
	for _, b := range bills {
		if b == 10 {
			if count[five] == 0 {
				return false
			}
			count[five]--
			count[ten]++
		} else if b == 20 {
			if count[ten] >= 1 && count[five] >= 1 {
				count[ten]--
				count[five]--
			} else if count[five] >= 3 {
				count[five] -= 3
			} else {
				return false
			}
		} else {
			count[five]++
		}
	}
	return true
}
