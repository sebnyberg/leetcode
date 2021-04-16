package removealladjacentduplicatesinstring2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"deeedbbcccbdaa", 3, "aa"},
		{"abcd", 2, "abcd"},
		{"pbbcggttciiippooaais", 2, "ps"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, removeDuplicates(tc.s, tc.k))
		})
	}
}

func removeDuplicates(s string, k int) string {
	type charCount struct {
		ch    rune
		count int
	}
	stack := make([]charCount, 0)
	n := 0
	for _, ch := range s {
		if len(stack) > 0 && stack[n-1].ch == ch {
			stack[n-1].count++
		} else {
			stack = append(stack, charCount{ch, 1})
			n++
		}
		if stack[n-1].count == k {
			// pop
			stack = stack[:n-1]
			n--
		}
	}

	var sb strings.Builder
	for _, c := range stack {
		for i := 0; i < c.count; i++ {
			sb.WriteRune(c.ch)
		}
	}
	return sb.String()
}
