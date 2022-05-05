package p0791customsortstring

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_customSortString(t *testing.T) {
	for _, tc := range []struct {
		order string
		str   string
		want  string
	}{
		{"kqep", "pekeq", "kepqe"},
		{"cba", "abcd", "cbad"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.order), func(t *testing.T) {
			require.Equal(t, tc.want, customSortString(tc.order, tc.str))
		})
	}
}

func customSortString(order string, str string) string {
	var charIdx [26]int
	for i := range charIdx {
		charIdx[i] = 30
	}
	for i, char := range order {
		charIdx[char-'a'] = i
	}
	bs := []byte(str)
	sort.Slice(bs, func(i, j int) bool {
		return charIdx[int(bs[i]-'a')] < charIdx[int(bs[j]-'a')]
	})
	res := string(bs)
	return res
}
