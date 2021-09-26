package daily_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxLength(t *testing.T) {
	for _, tc := range []struct {
		arr  []string
		want int
	}{
		{[]string{"un", "iq", "ue"}, 4},
		{[]string{"cha", "r", "act", "ers"}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, maxLength(tc.arr))
		})
	}
}

func maxLength(arr []string) int {
	res := make([]string, 0, len(arr)*2)
	var f uniqueFinder
	f.seen = make(map[string]struct{})
	f.seen[""] = struct{}{}
	f.dfs("", 0, len(arr), arr)
}

type uniqueFinder struct {
	seen map[string]struct{}
	res  []string
}

func (f *uniqueFinder) dfs(bm, idx, n int, arr []string) {
	if idx == n {
		if _, exists := pre
	}
}
