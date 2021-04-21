package p0267palindromeperm2

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test_generatePalindromes(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []string
	}{
		{"aaa", []string{"aaa"}},
		{"aabb", []string{"abba", "baab"}},
		{"abc", []string{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, generatePalindromes(tc.s))
		})
	}
}

func generatePalindromes(s string) []string {
	var chCount [26]int
	for _, ch := range s {
		chCount[ch-'a']++
	}
	oddIdx := -1
	for i, count := range chCount {
		if count%2 == 1 {
			if oddIdx != -1 {
				return []string{}
			}
			oddIdx = i
		}
		chCount[i] /= 2
	}

	var ntot int
	for _, count := range chCount {
		ntot += count
	}

	// form all left sides
	gen := &leftGenerator{
		lefts: make([][]byte, 0),
	}
	gen.generate(chCount, make([]byte, 0, 16), ntot)
	res := make([]string, len(gen.lefts))
	for i, left := range gen.lefts {
		n := len(left)
		if oddIdx != -1 {
			left = append(left, byte(oddIdx)+'a')
		}
		for j := n - 1; j >= 0; j-- {
			left = append(left, left[j])
		}
		res[i] = *(*string)(unsafe.Pointer(&left))
	}
	return res
}

type leftGenerator struct {
	lefts [][]byte
}

func (c *leftGenerator) generate(counts [26]int, prefix []byte, ntot int) {
	if len(prefix) == ntot {
		cpy := make([]byte, len(prefix))
		copy(cpy, prefix)
		c.lefts = append(c.lefts, cpy)
		return
	}
	n := len(prefix)
	for i, count := range counts {
		if count > 0 {
			counts[i]--
			prefix = append(prefix, byte(i)+'a')
			c.generate(counts, prefix, ntot)
			counts[i]++
			prefix = prefix[:n]
		}
	}
}
