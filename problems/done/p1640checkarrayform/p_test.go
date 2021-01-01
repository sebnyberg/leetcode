package p1640checkarrayform

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canFormArray(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		pieces [][]int
		want   bool
	}{
		// {[]int{85}, [][]int{{85}}, true},
		{[]int{15, 88}, [][]int{{88}, {15}}, true},
		// {[]int{49, 18, 16}, [][]int{{16, 18, 49}}, false},
		// {[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}, true},
		// {[]int{1, 3, 5, 7}, [][]int{{2, 4, 6, 8}}, false},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.arr, tc.pieces), func(t *testing.T) {
			require.Equal(t, tc.want, canFormArray(tc.arr, tc.pieces))
		})
	}
}

func canFormArray(arr []int, pieces [][]int) bool {
	// Since all integers are distinct, the first number of each piece is unique
	// Map the first number of each piece to its index in the pieces list
	pieceMap := make(map[int]int, len(pieces))
	for i, piece := range pieces {
		pieceMap[piece[0]] = i
	}

	for i := 0; i < len(arr); {
		first := arr[i]
		if _, exists := pieceMap[first]; !exists {
			return false
		}
		pieceIndex := pieceMap[first]
		for j := range pieces[pieceIndex] {
			if arr[i] != pieces[pieceIndex][j] {
				return false
			}
			i++
		}
		delete(pieceMap, first)
	}
	return true
}
