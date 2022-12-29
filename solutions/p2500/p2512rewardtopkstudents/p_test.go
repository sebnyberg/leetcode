package p2512rewardtopkstudents

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	scores := make(map[string]int)
	for _, f := range positive_feedback {
		scores[f] = 3
	}
	for _, f := range negative_feedback {
		scores[f] = -1
	}
	n := len(student_id)
	students := make([]student, n)
	for i := range students {
		students[i].idx = student_id[i]
		for _, w := range strings.Fields(report[i]) {
			students[i].score += scores[w]
		}
	}
	sort.Slice(students, func(i, j int) bool {
		if students[i].score == students[j].score {
			return students[i].idx < students[j].idx
		}
		return students[i].score > students[j].score
	})
	var res []int
	for _, s := range students {
		res = append(res, s.idx)
	}
	return res[:k]
}

type student struct {
	idx   int
	score int
}
