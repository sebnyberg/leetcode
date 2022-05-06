package p1209removealladjacentduplicatesinstring2

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
		{"abcd", 2, "abcd"},
		{"deeedbbcccbdaa", 3, "aa"},
		{"pbbcggttciiippooaais", 2, "ps"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, removeDuplicates(tc.s, tc.k))
		})
	}
}

func removeDuplicates(s string, k int) string {
	type charCount struct {
		char  rune
		count int
	}
	stack := []charCount{
		{0, 0},
	}
	for _, ch := range s {
		if ch != stack[len(stack)-1].char {
			stack = append(stack, charCount{ch, 1})
			continue
		}
		stack[len(stack)-1].count++
		for stack[len(stack)-1].count >= k {
			stack = stack[:len(stack)-1]
		}
	}
	var res []byte
	for _, cc := range stack[1:] {
		res = append(res, strings.Repeat(string(cc.char), cc.count)...)
	}
	return string(res)
}
