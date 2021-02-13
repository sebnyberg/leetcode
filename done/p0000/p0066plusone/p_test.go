package p0066plusone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_plusOne(t *testing.T) {
	for _, tc := range []struct {
		digits []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{0}, []int{1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.digits), func(t *testing.T) {
			require.Equal(t, tc.want, plusOne(tc.digits))
		})
	}
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carry == 0 {
			return digits
		}
		if digits[i] == 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i]++
			carry = 0
		}
	}
	if carry == 1 {
		digits = append(digits, 0)
		copy(digits[1:], digits)
		digits[0] = 1
	}
	return digits
}
