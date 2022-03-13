package p2201countartifactsthatcanbeextracted

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_digArtifacts(t *testing.T) {
	for i, tc := range []struct {
		n         int
		artifacts [][]int
		dig       [][]int
		want      int
	}{
		{
			2,
			leetcode.ParseMatrix("[[0,0,0,0],[0,1,1,1]]"),
			leetcode.ParseMatrix("[[0,0],[0,1],[1,1]]"),
			2,
		},
		{
			2,
			leetcode.ParseMatrix("[[0,0,0,0],[0,1,1,1]]"),
			leetcode.ParseMatrix("[[0,0],[0,1]]"),
			1,
		},
		{
			2,
			leetcode.ParseMatrix("[[0,0,0,0],[0,1,1,1]]"),
			leetcode.ParseMatrix("[[0,0],[0,1],[1,1]]"),
			2,
		},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			require.Equal(t, tc.want, digArtifacts(tc.n, tc.artifacts, tc.dig))
		})
	}
}

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	artifactCount := make([]int, len(artifacts)+1)
	var grid [1000][1000]int
	var seen [1000][1000]bool
	for i, artifact := range artifacts {
		for j := artifact[0]; j <= artifact[2]; j++ {
			for k := artifact[1]; k <= artifact[3]; k++ {
				grid[j][k] = i + 1
				artifactCount[i+1]++
			}
		}
	}
	var res int
	for _, d := range dig {
		if seen[d[0]][d[1]] {
			continue
		}
		i := grid[d[0]][d[1]]
		artifactCount[i]--
		if artifactCount[i] == 0 {
			res++
		}
		seen[d[0]][d[1]] = true
	}
	return res
}
