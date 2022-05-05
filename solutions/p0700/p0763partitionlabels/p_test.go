package p0763partitionlabels

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partitionLabels(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, partitionLabels(tc.s))
		})
	}
}

func partitionLabels(s string) []int {
	// There's a pretty simple, greedy solution for this exercise.
	// We find the right-most index of each letter
	var right [26]int
	for i := range s {
		right[s[i]-'a'] = i
	}

	var i int
	var res []int
	for i < len(s) {
		maxIdx := right[s[i]-'a']
		j := i + 1
		for ; j <= maxIdx; j++ {
			maxIdx = max(maxIdx, right[s[j]-'a'])
		}
		res = append(res, j-i)
		i = j
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
