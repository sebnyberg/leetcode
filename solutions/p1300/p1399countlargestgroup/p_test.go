package p1399countlargestgroup

import "fmt"

func countLargestGroup(n int) int {
	groups := map[int]int{}
	for x := 1; x <= n; x++ {
		var sum int
		for _, ch := range fmt.Sprint(x) {
			sum += int(ch - '0')
		}
		groups[sum]++
	}
	var largest int
	var largestGroups int
	for _, count := range groups {
		if count > largest {
			largest = count
			largestGroups = 1
		} else if count == largest {
			largestGroups++
		}
	}
	return largestGroups
}
