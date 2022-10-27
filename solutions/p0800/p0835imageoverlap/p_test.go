package p0835imageoverlap

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_largestOverlap(t *testing.T) {
	for _, tc := range []struct {
		img1 [][]int
		img2 [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,1],[1,1]]"),
			leetcode.ParseMatrix("[[1,1],[1,0]]"),
			2,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.img1), func(t *testing.T) {
			require.Equal(t, tc.want, largestOverlap(tc.img1, tc.img2))
		})
	}
}

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1[0])

	// Convert images to bitmaps
	var image1 [30]uint32
	var image2 [30]uint32
	for i := range img1 {
		for j := range img1 {
			image1[i] = (image1[i] << 1) | uint32(img1[i][j])
			image2[i] = (image2[i] << 1) | uint32(img2[i][j])
		}
	}

	// For each possible shift in x and y axis, try it and count common bits
	var res int
	for dy := -n + 1; dy < n; dy++ {
		for dx := 0; dx < n; dx++ {
			var left int
			var right int
			for i := 0; i < n; i++ {
				j := i + dy
				if j < 0 || j >= n {
					continue
				}
				// Note: we do not need to worry about going out of bounds
				// because the other image could not share those bits.
				left += bits.OnesCount32(image2[i] & (image1[j] << dx))
				right += bits.OnesCount32(image2[i] & (image1[j] >> dx))
			}
			res = max(res, left)
			res = max(res, right)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
