package p1189maxnumberofballoons

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNumberOfBalloons(t *testing.T) {
	for _, tc := range []struct {
		text string
		want int
	}{
		{"nlaebolko", 1},
		{"loonbalxballpoon", 2},
		{"leetcode", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.text), func(t *testing.T) {
			require.Equal(t, tc.want, maxNumberOfBalloons(tc.text))
		})
	}
}

func maxNumberOfBalloons(text string) int {
	var charCount [26]int
	for _, ch := range text {
		charCount[ch-'a']++
	}
	res := math.MaxInt32
	for _, ch := range [][2]int{
		{'b', 1},
		{'a', 1},
		{'l', 2},
		{'o', 2},
		{'n', 1},
	} {
		res = min(res, charCount[ch[0]-'a']/ch[1])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
