package p1901findapeakelement2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPeakGrid(t *testing.T) {
	for _, tc := range []struct {
		mat  [][]int
		want []int
	}{
		{
			[][]int{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{2, 3, 4, 5, 6, 7, 8, 9},
				{3, 4, 5, 6, 7, 8, 9, 10},
				{4, 5, 6, 7, 8, 9, 10, 11},
			},
			[]int{3, 7},
		},
		{[][]int{
			{47, 30, 35, 8, 25},
			{6, 36, 19, 41, 40},
			{24, 37, 13, 46, 5},
			{3, 43, 15, 50, 19},
			{6, 15, 7, 25, 18},
		},
			[]int{3, 3},
		},
		{[][]int{{1, 4}, {3, 2}}, []int{1, 0}},
		{[][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}, []int{1, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, findPeakGrid(tc.mat))
		})
	}
}

// func findPeakGrid(mat [][]int) []int {
// 	n := len(mat)
// 	l, r := 0, n

// 	for l < r {
// 		mid := (l + r) / 2
// 		var maxNum, maxIdx int
// 		for i, num := range mat[mid] {
// 			if num > maxNum {
// 				maxIdx = i
// 				maxNum = num
// 			}
// 		}
// 		switch {
// 		case mid > 0 && mat[mid-1][maxIdx] > maxNum:
// 			r = mid
// 		case mid < n-1 && mat[mid+1][maxIdx] > maxNum:
// 			l = mid + 1
// 		default:
// 			return []int{mid, maxIdx}
// 		}
// 	}
// 	return nil
// }

func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	var smallmat [503][503]uint16
	for i := range mat {
		for j := range mat[i] {
			smallmat[i+1][j+1] = uint16(mat[i][j])
		}
	}

	l, r := 1, m+1
	for l < r {
		mid := (l + r) / 2
		var maxNum uint16
		var maxIdx int
		for i, num := range smallmat[mid] {
			if num > maxNum {
				maxIdx = i
				maxNum = num
			}
		}
		switch {
		case smallmat[mid-1][maxIdx] > maxNum:
			r = mid
		case smallmat[mid+1][maxIdx] > maxNum:
			l = mid + 1
		default:
			return []int{mid - 1, maxIdx - 1}
		}
	}
	return nil
}

// func findPeakGrid(mat [][]int) []int {
// 	m, n := len(mat), len(mat[0])
// 	smallmat := make([][]uint16, m+2)
// 	smallmat[0] = make([]uint16, n+2)
// 	for i := range mat {
// 		smallmat[i+1] = make([]uint16, n+2)
// 		for j := range mat[i] {
// 			smallmat[i+1][j+1] = uint16(mat[i][j])
// 		}
// 	}
// 	smallmat[m+1] = make([]uint16, n+2)

// 	l, r := 1, m+1
// 	for l < r {
// 		mid := (l + r) / 2
// 		var maxNum uint16
// 		var maxIdx int
// 		for i, num := range smallmat[mid] {
// 			if num > maxNum {
// 				maxIdx = i
// 				maxNum = num
// 			}
// 		}
// 		switch {
// 		case smallmat[mid-1][maxIdx] > maxNum:
// 			r = mid
// 		case smallmat[mid+1][maxIdx] > maxNum:
// 			l = mid + 1
// 		default:
// 			return []int{mid - 1, maxIdx - 1}
// 		}
// 	}
// 	return nil
// }
