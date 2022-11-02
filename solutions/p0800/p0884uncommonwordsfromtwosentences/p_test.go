package p0884uncommonwordsfromtwosentences

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uncommonFromSentences(t *testing.T) {
	for i, tc := range []struct {
		s1   string
		s2   string
		want []string
	}{
		{"this apple is sweet", "this apple is sour", []string{"sweet", "sour"}},
		{"apple apple", "banana", []string{"banana"}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, uncommonFromSentences(tc.s1, tc.s2))
		})
	}
}

func uncommonFromSentences(s1 string, s2 string) []string {
	m1 := make(map[string]int)
	for _, w := range strings.Fields(s1) {
		m1[w]++
	}
	m2 := make(map[string]int)
	for _, w := range strings.Fields(s2) {
		m2[w]++
	}
	var res []string
	for w, c := range m1 {
		if c == 1 && m2[w] == 0 {
			res = append(res, w)
		}
	}
	for w, c := range m2 {
		if c == 1 && m1[w] == 0 {
			res = append(res, w)
		}
	}
	return res
}
