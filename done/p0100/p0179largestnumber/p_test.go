package p0179largestnumber

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestNumber(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want string
	}{
		// {[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, largestNumber(tc.nums))
		})
	}
}

func largestNumber(nums []int) string {
	n := len(nums)
	numStrs := make([]string, n)
	allZeroes := true
	for i, a := range nums {
		if a != 0 {
			allZeroes = false
		}
		numStrs[i] = strconv.Itoa(a)
	}
	if allZeroes {
		return "0"
	}
	sort.Slice(numStrs, func(i, j int) bool {
		a, b := numStrs[i], numStrs[j]
		switch {
		case a+b > b+a:
			return true
		case a+b < b+a:
			return false
		}
		return true
	})
	return strings.Join(numStrs, "")
}
