package p2409countdaysspenttogether

import "fmt"

var daysPerMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	dayOfYear := func(s string) int {
		var month, day int
		fmt.Sscanf(s, "%02d-%02d", &month, &day)
		for x := 1; x < month; x++ {
			day += daysPerMonth[x-1]
		}
		return day
	}
	a1, a2 := dayOfYear(arriveAlice), dayOfYear(leaveAlice)
	b1, b2 := dayOfYear(arriveBob), dayOfYear(leaveBob)
	if a1 > b1 {
		a1, a2, b1, b2 = b1, b2, a1, a2
	}
	overlap := a2 - b1 + 1
	if overlap <= 0 {
		return 0
	}
	if b2 < a2 {
		overlap -= a2 - b2
	}
	return overlap
}
