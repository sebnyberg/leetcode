package p2191

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortJumbled(t *testing.T) {
	for _, tc := range []struct {
		mapping []int
		nums    []int
		want    []int
	}{
		{[]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}, []int{338, 38, 991}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}, []int{123, 456, 789}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mapping), func(t *testing.T) {
			require.Equal(t, tc.want, sortJumbled(tc.mapping, tc.nums))
		})
	}
}

func sortJumbled(mapping []int, nums []int) []int {
	mapped := make([]mappedNum, len(nums))
	for i, num := range nums {
		mapped[i].idx = i
		s := []byte(fmt.Sprint(num))
		for j := range s {
			s[j] = byte(mapping[s[j]-'0'])
		}
		var x int
		for _, ch := range s {
			x = x*10 + int(ch)
		}
		mapped[i] = mappedNum{v: num, mapped: x}
	}
	sort.SliceStable(mapped, func(i, j int) bool {
		return mapped[i].mapped < mapped[j].mapped
	})
	res := make([]int, len(nums))
	for i := range mapped {
		res[i] = mapped[i].v
	}
	return res
}

type mappedNum struct {
	v      int
	mapped int
	idx    int
}
