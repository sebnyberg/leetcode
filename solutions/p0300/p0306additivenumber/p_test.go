package p0306additivenumber

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isAdditiveNumber(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want bool
	}{
		{"101", true},
		{"000", true},
		{"11000000000010000000001", true},
		{"11011", true},
		{"112358", true},
		{"199100199", true},
		{"00111", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, isAdditiveNumber(tc.num))
		})
	}
}

func isAdditiveNumber(num string) bool {
	// Since num is <= 35 characters long, the biggest possible additive number
	// is less in length than an int64 (20 characters). We can therefore use
	// ints to parse num, even though it is less efficient than accepting
	// runes instead.
	n := len(num)
	for i := 1; i <= n/2; i++ {
		first, _ := strconv.Atoi(num[:i])
		for j := 1; j+i < n; j++ {
			second, _ := strconv.Atoi(num[i : i+j])
			if helper(num, first, second, i+j) {
				return true
			}
			// If the current number is a zero, it is the only acceptable value
			if num[i] == '0' {
				break
			}
		}
		if num[0] == '0' {
			break
		}
	}
	return false
}

func helper(num string, first, second, pos int) bool {
	for i := 1; pos+i <= len(num) && i <= len(num)/2; i++ {
		newSecond, err := strconv.Atoi(num[pos : pos+i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		if first+second == newSecond {
			if pos+i == len(num) || helper(num, second, newSecond, pos+i) {
				return true
			}
		}
		if num[pos] == '0' {
			break
		}
	}
	return false
}
