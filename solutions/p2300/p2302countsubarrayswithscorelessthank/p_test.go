package p2302countsubarrayswithscorelessthank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matchReplacement(t *testing.T) {
	for _, tc := range []struct {
		s        string
		sub      string
		mappings [][]byte
		want     bool
	}{
		{"fool3e7bar", "leet", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}}, true},
		{"fooleetbar", "f00l", [][]byte{{'o', '0'}}, false},
		{"Fool33tbaR", "leetd", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}, {'d', 'b'}, {'p', 'b'}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, matchReplacement(tc.s, tc.sub, tc.mappings))
		})
	}
}

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	var mappable [256][256]bool
	for i := 0; i < 256; i++ {
		mappable[i][i] = true
	}
	for _, m := range mappings {
		mappable[m[0]][m[1]] = true
	}
	check := func(s string) bool {
		for i := range sub {
			if !mappable[sub[i]][s[i]] {
				return false
			}
		}
		return true
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if check(s[i : i+len(sub)]) {
			return true
		}
	}
	return false
}
