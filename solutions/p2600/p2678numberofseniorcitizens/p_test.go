package p2678numberofseniorcitizens

import "strconv"

func countSeniors(details []string) int {
	var res int
	for _, s := range details {
		age, _ := strconv.Atoi(s[11:13])
		if age > 60 {
			res++
		}
	}
	return res
}
