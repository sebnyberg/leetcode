package p1700numberofstudentsunabletoeatlunch

func countStudents(students []int, sandwiches []int) int {
	var count [2]int
	for _, s := range students {
		count[s]++
	}
	res := len(sandwiches)
	for _, s := range sandwiches {
		if count[s] == 0 {
			break
		}
		count[s]--
		res--
	}
	return res
}
