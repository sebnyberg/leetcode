package p2446determineiftwoeventshaveconflict

import "fmt"

func haveConflict(event1 []string, event2 []string) bool {

	parse := func(d string) int {
		var hh, mm int
		fmt.Sscanf(d, "%d:%d", &hh, &mm)
		return hh*60 + mm
	}
	s1, e1 := parse(event1[0]), parse(event1[1])
	s2, e2 := parse(event2[0]), parse(event2[1])
	if s1 > s2 {
		s1, s2 = s2, s1
		e1, e2 = e2, e1
	}
	return s2 <= e1
}
