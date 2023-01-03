package p1073addingtwonegabinarynumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addNegabinary(t *testing.T) {
	for i, tc := range []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{[]int{1, 0, 1, 1}, []int{1, 1, 1, 0}, []int{1, 1, 0, 0, 0, 1}},
		{[]int{1, 0, 1, 0, 1, 0}, []int{1, 0, 1, 1, 0, 0}, []int{1, 1, 1, 1, 0, 1, 1, 0}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, addNegabinary(tc.arr1, tc.arr2))
		})
	}
}

func addNegabinary(arr1 []int, arr2 []int) []int {
	rev := func(a []int) []int {
		for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
			a[l], a[r] = a[r], a[l]
		}
		return a
	}
	arr1 = rev(arr1)
	arr2 = rev(arr2)
	arr1 = append(arr1, make([]int, max(0, len(arr2)-len(arr1)))...)
	arr2 = append(arr2, make([]int, max(0, len(arr1)-len(arr2)))...)
	var res []int
	var posCarry int
	var negCarry int
	for i := range arr1 {
		x := arr1[i] + arr2[i] + negCarry + posCarry
		if x < 0 {
			res = append(res, 1)
			posCarry = 1
			negCarry = 0
			continue
		}
		if x >= 2 {
			negCarry = -1
			posCarry = 0
		} else {
			negCarry = 0
			posCarry = 0
		}
		res = append(res, x&1)
	}
	carry := posCarry + negCarry
	if carry == -1 {
		res = append(res, 1, 1)
	} else if carry == 1 {
		res = append(res, 1)
	}
	// Trim any trailing zeroes
	for i := len(res) - 1; i >= 1 && res[i] == 0; i-- {
		res = res[:i]
	}
	ret := rev(res)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
