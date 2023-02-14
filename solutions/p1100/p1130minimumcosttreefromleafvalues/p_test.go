package p1130minimumcosttreefromleafvalues

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mctFromLeafValues(t *testing.T) {
	for i, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{11, 12, 12}, 276},
		{[]int{6, 2, 4}, 32},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, mctFromLeafValues(tc.arr))
		})
	}
}

func mctFromLeafValues(arr []int) int {
	// Given each level of the tree, we choose where to cut the tree.
	//
	// The total cost of the tree will be max(left)*max(right) + cost(left) +
	// cost(right). The cost of a single-leaf tree is zero.
	//
	mem := make(map[[2]int]int)
	res := dp(mem, arr, 0, len(arr))
	return res
}

func dp(mem map[[2]int]int, arr []int, l, r int) int {
	n := r - l
	if n <= 1 {
		return 0
	}
	k := [2]int{l, r}
	if v, exists := mem[k]; exists {
		return v
	}
	maxLeft := arr[l]
	res := math.MaxInt32
	for i := l + 1; i < r; i++ {
		maxRight := arr[i]
		for j := i + 1; j < r; j++ {
			maxRight = max(maxRight, arr[j])
		}
		res = min(res, dp(mem, arr, l, i)+dp(mem, arr, i, r)+maxLeft*maxRight)
		maxLeft = max(maxLeft, arr[i])
	}
	mem[k] = res
	return res
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
