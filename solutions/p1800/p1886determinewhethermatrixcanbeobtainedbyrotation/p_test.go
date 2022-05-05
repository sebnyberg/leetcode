package p1886determinewhethermatrixcanbeobtainedbyrotation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRotation(t *testing.T) {
	for _, tc := range []struct {
		mat    [][]int
		target [][]int
		want   bool
	}{
		{[][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}, true},
		{[][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}, false},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, findRotation(tc.mat, tc.target))
		})
	}
}

func findRotation(mat [][]int, target [][]int) bool {
	if ok(mat, target) {
		return true
	}
	n := len(mat)
	for i := 0; i < 3; i++ {
		// Rotate
		newMat := make([][]int, n)
		for i := range mat {
			newMat[i] = make([]int, n)
		}
		for i := range mat {
			col := n - i - 1
			for j := range mat[0] {
				newMat[j][col] = mat[i][j]
			}
		}
		if ok(newMat, target) {
			return true
		}
		mat = newMat
	}
	return false
}

func ok(a, b [][]int) bool {
	for i := range a {
		for j := range a[0] {
			if b[i][j] != a[i][j] {
				return false
			}
		}
	}
	return true
}
