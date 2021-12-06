package p2094finding3digitevennumbers

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findEvenNumbers(t *testing.T) {
	for _, tc := range []struct {
		digits []int
		want   []int
	}{
		{[]int{2, 1, 3, 0}, []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}},
		{[]int{2, 2, 8, 8, 2}, []int{222, 228, 282, 288, 822, 828, 882}},
		{[]int{3, 7, 5}, []int{}},
		{[]int{0, 2, 0, 0}, []int{200}},
		{[]int{0, 0, 0}, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.digits), func(t *testing.T) {
			require.Equal(t, tc.want, findEvenNumbers(tc.digits))
		})
	}
}

func findEvenNumbers(digits []int) []int {
	uniq := make(map[int]struct{})
	add := func(val int) {
		if val >= 100 && val%2 == 0 {
			uniq[val] = struct{}{}
		}
	}
	for i := 0; i < len(digits)-2; i++ {
		for j := i + 1; j < len(digits)-1; j++ {
			for k := j + 1; k < len(digits); k++ {
				add(100*digits[i] + 10*digits[j] + digits[k])
				add(100*digits[i] + 10*digits[k] + digits[j])
				add(100*digits[j] + 10*digits[i] + digits[k])
				add(100*digits[j] + 10*digits[k] + digits[i])
				add(100*digits[k] + 10*digits[i] + digits[j])
				add(100*digits[k] + 10*digits[j] + digits[i])
			}
		}
	}
	vals := make([]int, 0, len(uniq))
	for val := range uniq {
		vals = append(vals, val)
	}
	sort.Ints(vals)
	return vals
}
