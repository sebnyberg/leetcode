package p1980finduniquebinarystring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findUniqueBinaryString(t *testing.T) {
	for _, tc := range []struct {
		nums []string
		want string
	}{
		{[]string{"01", "10"}, "11"},
		{[]string{"00", "01"}, "11"},
		{[]string{"111", "011", "001"}, "101"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			exists := make(map[string]bool)
			for _, num := range tc.nums {
				exists[num] = true
			}
			res := findDifferentBinaryString(tc.nums)
			require.False(t, exists[res])
		})
	}
}

func findDifferentBinaryString(nums []string) string {
	// There are up to 16 elements, and up to 16 bits per element,
	// An easy solution is to simply test all possible elements...
	var seen [1 << 16]bool
	for _, num := range nums {
		var val int
		for _, ch := range num {
			val <<= 1
			if ch == '1' {
				val++
			}
		}
		seen[val] = true
	}
	var match int
	for seen[match] {
		match++
	}
	n := len(nums[0])
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[n-1-i] = byte(match&1) + '0'
		match >>= 1
	}
	return string(res)
}
