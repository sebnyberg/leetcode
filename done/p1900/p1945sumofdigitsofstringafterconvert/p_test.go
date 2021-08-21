package p1945sumofdigitsofstringafterconvert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getLucky(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"leetcode", 2, 6},
		{"iiii", 1, 36},
		{"zbax", 2, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, getLucky(tc.s, tc.k))
		})
	}
}

func getLucky(s string, k int) int {
	// Transform string to sequence of ints
	var sum int
	for i := range s {
		num := int(s[i] - 'a' + 1)
		for num > 0 {
			sum += num % 10
			num /= 10
		}
	}

	nums := make([]int, 0, len(s))
	for i := 1; i < k; i++ {
		nums = nums[:0]
		for sum > 0 {
			nums = append(nums, sum%10)
			sum /= 10
		}
		for _, num := range nums {
			sum += num
		}
	}

	return sum
}
