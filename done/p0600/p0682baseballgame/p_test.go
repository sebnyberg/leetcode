package p0682baseballgame

import (
	"log"
	"strconv"
)

func calPoints(ops []string) int {
	var scores []int
	for _, op := range ops {
		n := len(scores)
		switch op {
		case "C":
			scores = scores[:n-1]
		case "D":
			scores = append(scores, 2*scores[n-1])
		case "+":
			scores = append(scores, scores[n-2]+scores[n-1])
		default: // integer
			x, err := strconv.Atoi(op)
			if err != nil {
				log.Fatalln(err)
			}
			scores = append(scores, x)
		}
	}
	var sum int
	for _, score := range scores {
		sum += score
	}
	return sum
}
