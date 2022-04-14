package p0661imagesmoother

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_imageSmoother(t *testing.T) {
	type testCase struct {
		img  [][]int
		want [][]int
	}

	testCases := []testCase{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{
			[][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
			[][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase %v", i), func(t *testing.T) {
			require.Equal(t, tc.want, imageSmoother(tc.img))
		})
	}
}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	for i := range img {
		for j := range img[i] {
			sum := img[i][j]
			count := 1
			for _, d := range dirs {
				ii, jj := i+d[0], j+d[1]
				if ok(ii, jj) {
					sum += img[ii][jj]
					count++
				}
			}
			res[i][j] = sum / count
		}
	}
	return res
}
