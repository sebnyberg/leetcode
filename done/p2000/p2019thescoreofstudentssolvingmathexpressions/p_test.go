package p2019thescoreofstudentssolvingmathexpressions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_scoreOfStudents(t *testing.T) {
	for _, tc := range []struct {
		s       string
		answers []int
		want    int
	}{
		{"7+3*1*2", []int{20, 13, 42}, 7},
		{"3+5*2", []int{13, 0, 10, 13, 13, 16, 16}, 19},
		{"6+0*1", []int{12, 9, 6, 4, 8, 6}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, scoreOfStudents(tc.s, tc.answers))
		})
	}
}

func scoreOfStudents(s string, answers []int) int {
	n := len(s)
	numVals := (n / 2) + 1
	dp := make([][]map[int]struct{}, numVals)
	for i := range dp {
		dp[i] = make([]map[int]struct{}, numVals)
		for j := range dp[i] {
			dp[i][j] = make(map[int]struct{})
		}
	}
	for i := 0; i < numVals; i++ {
		dp[i][i][int(s[2*i]-'0')] = struct{}{}
	}
	// Calculate all possible values for expressions within a certain window
	// length.
	for width := 2; width <= numVals; width++ {
		for start := 0; start <= numVals-width; start++ {
			end := start + width - 1
			for i := start*2 + 1; i < end*2+1; i += 2 {
				for a := range dp[start][i/2] {
					for b := range dp[i/2+1][end] {
						var val int
						if s[i] == '+' {
							val = a + b
						} else {
							val = a * b
						}
						if val <= 1000 {
							dp[start][end][val] = struct{}{}
						}
					}
				}
			}
		}
	}
	var result int
	correct := eval(s)
	for _, ans := range answers {
		if ans == correct {
			result += 5
			continue
		}
		if _, exists := dp[0][numVals-1][ans]; exists {
			result += 2
		}
	}
	return result
}

func eval(s string) int {
	vals := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			vals[len(vals)-1] *= int(s[i+1] - '0')
			i++
		} else if s[i] != '+' {
			vals = append(vals, int(s[i]-'0'))
		}
	}
	var res int
	for _, v := range vals {
		res += v
	}
	return res
}
