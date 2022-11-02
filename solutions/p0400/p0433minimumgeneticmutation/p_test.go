package p0433minimumgeneticmutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMutation(t *testing.T) {
	for _, tc := range []struct {
		start string
		end   string
		bank  []string
		want  int
	}{
		{"AGCAAAAA", "GACAAAAA", []string{"AGTAAAAA", "GGTAAAAA", "GATAAAAA", "GACAAAAA"}, 4},
		{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.start), func(t *testing.T) {
			require.Equal(t, tc.want, minMutation(tc.start, tc.end, tc.bank))
		})
	}
}

func minMutation(start string, end string, bank []string) int {
	dist := func(a, b string) int {
		var d int
		for i := range a {
			if a[i] != b[i] {
				d++
			}
		}
		return d
	}
	curr := []string{start}
	next := []string{}
	seen := make(map[string]struct{}, len(bank))
	seen[start] = struct{}{}
	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, a := range curr {
			for _, b := range bank {
				_, exists := seen[b]
				if exists || dist(a, b) > 1 {
					continue
				}
				if b == end {
					return steps
				}
				seen[b] = struct{}{}
				next = append(next, b)
			}
		}
		curr, next = next, curr
	}
	return -1
}
