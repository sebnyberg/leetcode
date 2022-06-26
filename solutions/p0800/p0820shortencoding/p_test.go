package p0820shortencoding

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumLengthEncoding(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"time", "atime", "btime"}, 12},
		{[]string{"time", "me", "bell"}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, minimumLengthEncoding(tc.words))
		})
	}
}

func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	var res int
	m := make(map[string]struct{})
	for _, w := range words {
		if _, exists := m[w]; exists {
			continue
		}
		res += len(w) + 1
		for j := 0; j < len(w); j++ {
			m[w[j:]] = struct{}{}
		}
	}
	return res
}
