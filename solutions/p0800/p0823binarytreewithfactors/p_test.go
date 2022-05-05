package p0823binarytreewithfactors

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numFactoredBinaryTrees(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		// {[]int{2, 4}, 3},
		// {[]int{2, 4, 5, 10}, 7},
		{[]int{46, 144, 5040, 4488, 544, 380, 4410, 34, 11, 5, 3063808, 5550, 34496, 12, 540, 28, 18, 13, 2, 1056, 32710656, 31, 91872, 23, 26, 240, 18720, 33, 49, 4, 38, 37, 1457, 3, 799, 557568, 32, 1400, 47, 10, 20774, 1296, 9, 21, 92928, 8704, 29, 2162, 22, 1883700, 49588, 1078, 36, 44, 352, 546, 19, 523370496, 476, 24, 6000, 42, 30, 8, 16262400, 61600, 41, 24150, 1968, 7056, 7, 35, 16, 87, 20, 2730, 11616, 10912, 690, 150, 25, 6, 14, 1689120, 43, 3128, 27, 197472, 45, 15, 585, 21645, 39, 40, 2205, 17, 48, 136}, 509730797},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, numFactoredBinaryTrees(tc.arr))
		})
	}
}

const mod = int(1e9 + 7)

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	subTrees := make(map[int]int, len(arr))
	var res int
	for i, a := range arr {
		subTrees[a] = 1
		for j := i - 1; j >= 0; j-- {
			if a%arr[j] != 0 {
				continue
			}
			v := a / arr[j]
			if n, exists := subTrees[v]; exists {
				subTrees[a] += n * subTrees[arr[j]]
			}
		}
		res += subTrees[a] % mod
	}

	return res
}
