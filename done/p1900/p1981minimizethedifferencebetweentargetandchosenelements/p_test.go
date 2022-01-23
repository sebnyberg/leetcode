package p1981minimizethedifferencebetweentargetandchosenelements

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func ParseMatrix(s string) [][]int {
	s = s[2 : len(s)-2]
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "],[")
	res := make([][]int, len(parts))
	for i, part := range parts {
		if part == "" {
			continue
		}
		for _, numStr := range strings.Split(part, ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("failed to parse number, %v, %v\n", numStr, err)
			}
			res[i] = append(res[i], num)
		}
	}
	return res
}

func Test_minimizeTheDifference(t *testing.T) {
	for _, tc := range []struct {
		mat    [][]int
		target int
		want   int
	}{
		{ParseMatrix("[[3,5],[5,10]]"), 47, 32},
		{ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]"), 13, 0},
		{ParseMatrix("[[1],[2],[3]]"), 100, 94},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, minimizeTheDifference(tc.mat, tc.target))
		})
	}
}

func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])

	// Find min/max possible sum per row
	minPossible := make([]int, m)
	maxPossible := make([]int, m)
	for i := range mat {
		minVal := math.MaxInt32
		var maxVal int
		for j := range mat[i] {
			minVal = min(minVal, mat[i][j])
			maxVal = max(maxVal, mat[i][j])
		}
		if i > 0 {
			minPossible[i] = minPossible[i-1] + minVal
			maxPossible[i] = maxPossible[i-1] + maxVal
		} else {
			minPossible[i] = minVal
			maxPossible[i] = maxVal
		}
	}

	// The numbers are small (<=70) and the size of the grid as well (n<=70)
	// The number of possible sums is just 4900
	const size = 70*70 + 1
	var prev [size]bool
	for _, n := range mat[0] {
		prev[n] = true
	}

	// Find possible sums
	for i := 1; i < m; i++ {
		var cur [size]bool
		for sum := minPossible[i-1]; sum <= maxPossible[i-1]; sum++ {
			if prev[sum] {
				for j := 0; j < n; j++ {
					cur[sum+mat[i][j]] = true
				}
			}
		}
		prev = cur
	}

	// Find minimum delta
	minDelta := math.MaxInt32
	for sum := 0; sum < len(prev); sum++ {
		if !prev[sum] {
			continue
		}
		if d := abs(sum - target); d < minDelta {
			minDelta = d
		}
	}
	return minDelta
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
