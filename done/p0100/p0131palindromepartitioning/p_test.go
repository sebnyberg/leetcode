package p0131palindromepartitioning

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, partition(tc.s))
		})
	}
}

func partition(s string) [][]string {
	var f partitionFinder
	f.partitions = make([][]int, 0)
	f.findPartitions(s, []int{})
	res := make([][]string, len(f.partitions))
	for i, partition := range f.partitions {
		var pos int
		for _, cut := range partition {
			res[i] = append(res[i], s[pos:pos+cut])
			pos += cut
		}
	}
	return res
}

type partitionFinder struct {
	partitions [][]int
}

func (f *partitionFinder) findPartitions(s string, prefix []int) {
	if len(s) == 0 {
		f.partitions = append(f.partitions, prefix)
		return
	}
	var didCopy bool
	n := len(prefix)
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			if !didCopy {
				prefix = append(prefix, i)
				f.findPartitions(s[i:], prefix)
				didCopy = true
			} else {
				prefixCpy := make([]int, n)
				copy(prefixCpy, prefix[:n])
				prefixCpy = append(prefixCpy, i)
				f.findPartitions(s[i:], prefixCpy)
			}
		}
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < (len(s)+1)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
