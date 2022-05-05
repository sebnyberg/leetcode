package p1947maximumcompatibilityscoresum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCompatibilitySum(t *testing.T) {
	for _, tc := range []struct {
		students, mentors [][]int
		want              int
	}{
		{[][]int{{0, 1, 0, 1, 1, 1}, {1, 0, 0, 1, 0, 1}, {1, 0, 1, 1, 0, 0}}, [][]int{{1, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 1, 1}, {0, 1, 0, 0, 1, 1}}, 10},
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}, 8},
		{[][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.students), func(t *testing.T) {
			require.Equal(t, tc.want, maxCompatibilitySum(tc.students, tc.mentors))
		})
	}
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	// There are 8! different pairs, which is a manageable number.
	// If we compute the matching between all students and mentors, it should
	// be possible to find a solution through DFS.
	var compat [9][9]int
	for i := range students {
		for j := range mentors {
			compat[i][j] = calcCompatibility(students[i], mentors[j])
		}
	}

	f := maxCompatFinder{
		maxCompat: 0,
		compat:    compat,
	}

	// Try all possibilities
	f.findMaxCompat(0, 0, 0, len(students))
	return f.maxCompat
}

type maxCompatFinder struct {
	maxCompat int
	compat    [9][9]int
}

func (f *maxCompatFinder) findMaxCompat(mentorAvailability, curCompat, studentIdx, maxIdx int) {
	if studentIdx == maxIdx {
		if curCompat > f.maxCompat {
			f.maxCompat = curCompat
		}
		// f.maxCompat = max(f.maxCompat, curCompat)
		return
	}
	// Try each mentor
	for i := 0; i < maxIdx; i++ {
		b := 1 << i
		if mentorAvailability&b > 0 {
			continue
		}
		f.findMaxCompat(mentorAvailability|b, curCompat+f.compat[studentIdx][i], studentIdx+1, maxIdx)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcCompatibility(student, mentor []int) int {
	var res int
	for i := 0; i < len(student); i++ {
		if student[i] == mentor[i] {
			res++
		}
	}
	return res
}
