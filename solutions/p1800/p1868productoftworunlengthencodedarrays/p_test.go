package p1868productoftworunlengthencodedarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRLEArray(t *testing.T) {
	for _, tc := range []struct {
		encoded1 [][]int
		encoded2 [][]int
		want     [][]int
	}{
		{[][]int{{1, 3}, {2, 1}, {3, 2}}, [][]int{{2, 3}, {3, 3}}, [][]int{{2, 3}, {6, 1}, {9, 2}}},
		{[][]int{{1, 3}, {2, 3}}, [][]int{{6, 3}, {3, 3}}, [][]int{{6, 6}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.encoded1), func(t *testing.T) {
			require.Equal(t, tc.want, findRLEArray(tc.encoded1, tc.encoded2))
		})
	}
}

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	// encoded[i] = [val_i, freq_i]
	curNum := -10001
	curCount := 0
	var i, j int
	res := make([][]int, 0)
	for i < len(encoded1) && j < len(encoded2) {
		prod := encoded1[i][0] * encoded2[j][0]
		if prod != curNum {
			if curNum != -10001 {
				res = append(res, []int{curNum, curCount})
				curCount = 0
			}
			curNum = prod
		}
		l1, l2 := encoded1[i][1], encoded2[j][1]
		switch {
		case l1 == l2:
			i++
			j++
			curCount += l1
		case l1 < l2:
			i++
			encoded2[j][1] -= l1
			curCount += l1
		case l1 > l2:
			j++
			encoded1[i][1] -= l2
			curCount += l2
		}
	}
	res = append(res, []int{curNum, curCount})
	return res
}
