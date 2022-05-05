package p2260minimumconsecutivecardstopickup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_appealSum(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int64
	}{
		{"abbca", 28},
		{"code", 20},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, appealSum(tc.s))
		})
	}
}

func appealSum(s string) int64 {
	// Consider each letter,
	// consider all substrings that end at that letter,
	// Any substring which contains a certain letter will increase the total
	// appeal by 1 per substring
	// And so we record the first position of each letter so that we can count how
	// much appeal is added for each position
	var lastPos [26]int
	var res int
	for i, ch := range s {
		lastPos[ch-'a'] = i + 1
		for _, j := range lastPos {
			res += j
		}
	}
	return int64(res)
}
