package p0167twosum2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	for _, tc := range []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.numbers, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, twoSum(tc.numbers, tc.target))
		})
	}
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		switch {
		case numbers[l]+numbers[r] > target:
			r--
		case numbers[l]+numbers[r] == target:
			return []int{l + 1, r + 1}
		case numbers[l]+numbers[r] < target:
			l++
		}
	}
	return nil
}
