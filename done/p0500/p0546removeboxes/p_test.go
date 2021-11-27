package p0546removeboxes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeBoxes(t *testing.T) {
	for _, tc := range []struct {
		boxes []int
		want  int
	}{
		{[]int{1, 3, 2, 2, 2, 3, 4, 3, 1}, 23},
		{[]int{1, 1, 1}, 9},
		{[]int{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.boxes), func(t *testing.T) {
			require.Equal(t, tc.want, removeBoxes(tc.boxes))
		})
	}
}

func removeBoxes(boxes []int) int {
	n := len(boxes)
	var dp [100][100][100]int
	for j := 0; j < n; j++ {
		for k := 0; k <= j; k++ {
			dp[j][j][k] = (k + 1) * (k + 1)
		}
	}

	for l := 1; l < n; l++ {
		for j := l; j < n; j++ {
			i := j - l
			for k := 0; k <= i; k++ {
				res := (k+1)*(k+1) + dp[i+1][j][0]
				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[i] {
						res = max(res, dp[i+1][m-1][0]+dp[m][j][k+1])
					}
				}
				dp[i][j][k] = res
			}
		}
	}

	if n == 0 {
		return 0
	}
	return dp[0][n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
