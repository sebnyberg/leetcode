package p0049groupanagrams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[[32]byte][]string)
	for _, s := range strs {
		var count [32]byte
		for _, ch := range s {
			count[ch-'a']++
		}
		m[count] = append(m[count], s)
	}
	res := make([][]string, 0, len(m))
	for _, ss := range m {
		res = append(res, ss)
	}
	return res
}
