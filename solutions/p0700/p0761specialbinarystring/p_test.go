package p0761specialbinarystring

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeLargestSpecial(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"1010101100", "1100101010"},
		{"101010", "101010"},
		{"11011000", "11100100"},
		{"10", "10"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, makeLargestSpecial(tc.s))
		})
	}
}

func makeLargestSpecial(s string) string {
	var l int
	s = "1" + s + "0"
	stack := [][]string{{}}
	for _, ch := range s {
		if ch == '1' {
			stack[l] = append(stack[l], "1")
			l++
			if len(stack) <= l {
				stack = append(stack, []string{})
			}
			continue
		}

		// '0' => sort, concatenate, clean, and combine with lower level
		sort.Slice(stack[l], func(i, j int) bool {
			return stack[l][i] > stack[l][j]
		})

		ss := strings.Join(stack[l], "")
		stack[l] = stack[l][:0]
		l--
		n := len(stack[l])
		stack[l][n-1] = stack[l][n-1] + ss + "0"
	}

	return stack[0][0][1 : len(s)-1]
}
