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
	// A mutation can be considered a replacement of a character such that the
	// result belongs in the bank.

	// For each gene, create "mappers" where one letter has been removed from the
	// gene. If two genes could end on the same mapper, then they must be
	// compatible.
	findMappers := func(s string) []string {
		bs := []byte(s)
		res := make([]string, 0, len(s)-1)
		for i := range bs {
			res = append(res, string(bs[:i])+string('*')+string(bs[i+1:]))
		}
		return res
	}

	// End must be in the gene bank or there is no solution to this problem
	var hasEnd bool
	for _, gene := range bank {
		if gene == end {
			hasEnd = true
			break
		}
	}
	if !hasEnd {
		return -1
	}

	mapperGenes := make(map[string][]string)
	for _, gene := range bank {
		if gene == end || gene == start { // don't add start / end twice..
			continue
		}
		for _, m := range findMappers(gene) {
			mapperGenes[m] = append(mapperGenes[m], gene)
		}
	}
	for _, m := range findMappers(end) {
		mapperGenes[m] = append(mapperGenes[m], end)
	}

	seen := make(map[string]struct{})
	seen[start] = struct{}{}
	cur := []string{start}
	next := []string{}
	var mutations int
	for len(cur) > 0 {
		mutations++
		next = next[:0]
		for _, s := range cur {
			// Find possible mutations
			for _, mapper := range findMappers(s) {
				for _, gene := range mapperGenes[mapper] {
					if _, exists := seen[gene]; exists {
						continue
					}
					seen[gene] = struct{}{}
					if gene == end {
						return mutations
					}
					next = append(next, gene)
				}
			}
		}
		cur, next = next, cur
	}
	return -1
}
