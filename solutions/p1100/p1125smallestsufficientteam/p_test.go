package p1125smallestsufficientteam

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestSufficientTeam(t *testing.T) {
	for i, tc := range []struct {
		reqSkills []string
		people    [][]string
		want      []int
	}{
		{
			[]string{"java", "nodejs", "reactjs"},
			[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
			[]int{0, 2},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, smallestSufficientTeam(tc.reqSkills, tc.people))
		})
	}
}

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	// We want to find the smallest team which satisfies all required skills.
	//
	// This seems like an enumeration / bit problem due to the 16 possible
	// skills. Let's start by indexing skills.

	m := len(reqSkills)
	skillIdx := make(map[string]int)
	for i, s := range reqSkills {
		skillIdx[s] = i
	}

	n := len(people)
	peopleSkills := make([]int, n)
	for i := range people {
		var bm int
		for _, s := range people[i] {
			bm |= (1 << skillIdx[s])
		}
		peopleSkills[i] = bm
	}

	// The goal is to achieve a full bitmap, i.e. (1 << m) - 1
	want := (1 << m) - 1

	// The question is what pairing of people will achive that bitmap.
	// Since there can only be 16 possible skills, the maximum size is 2^16 ~=
	// 56k. Then, there are 60 people. If we combine each person with every
	// possible other prior skill, we get 60*56k ~= 2.5M ops which is definitely
	// fast enough.
	//

	curr := make([][]int, 1<<m)
	curr[0] = []int{-1}
	next := make([][]int, 1<<m)
	next[0] = []int{-1}

	for i := range people {
		bm := peopleSkills[i]
		for i := range next {
			next[i] = curr[i]
		}
		for x := range curr {
			if len(curr[x]) == 0 {
				continue
			}
			y := x | bm
			if len(next[y]) == 0 {
				next[y] = make([]int, len(curr[x])+1)
				copy(next[y], curr[x])
				next[y][len(next[y])-1] = i
				continue
			}

			if len(next[y]) <= len(curr[x])+1 {
				continue
			}

			// Move curr[x] to curr[y]
			next[y] = make([]int, len(curr[x])+1)
			copy(next[y], curr[x])
			next[y][len(next[y])-1] = i
		}
		curr, next = next, curr
	}

	return curr[want][1:]
}

func max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
