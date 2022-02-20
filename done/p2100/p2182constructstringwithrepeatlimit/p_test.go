package p2182constructstringwithrepeatlimit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_repeatLimitedString(t *testing.T) {
	for _, tc := range []struct {
		s           string
		repeatLimit int
		want        string
	}{
		{"cczazcc", 3, "zzcccac"},
		{"ccccccc", 3, "ccc"},
		{"aababab", 2, "bbabaa"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, repeatLimitedString(tc.s, tc.repeatLimit))
		})
	}
}

func repeatLimitedString(s string, repeatLimit int) string {
	// We can be greedy here, i.e. it's preferable to choose a high character
	// rather than low even if choosing it means a shorter result
	var charCount [26]int
	for _, ch := range s {
		charCount[ch-'a']++
	}
	var charDist [26]int
	for i := range charDist {
		charDist[i] = 0
	}
	res := make([]rune, 1, len(s))
	res[0] = 0
	for i := 1; ; {
		for ch := 'z'; ch >= 'a'; ch-- {
			if charCount[ch-'a'] == 0 || (res[i-1] == ch && charDist[ch-'a'] == repeatLimit) {
				continue
			}
			charCount[ch-'a']--
			charDist[ch-'a']++
			if i > 1 && res[i-1] != ch {
				charDist[res[i-1]-'a'] = 0
			}
			res = append(res, ch)
			i++
			goto continueLoop
		}
		break
	continueLoop:
	}

	ss := string(res[1:])
	return ss
}
