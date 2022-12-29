package p2515shortestdistancetotargetstringinacirculararray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closetTarget(t *testing.T) {
	for i, tc := range []struct {
		words      []string
		target     string
		startIndex int
		want       int
	}{
		{
			[]string{"odjrjznxpn", "cyulttuabe", "zqxkdoeszk", "yeewpgriok", "odjrjznxpn", "btqpvxpjzv", "ukyudladhk", "ukyudladhk", "odjrjznxpn", "yeewpgriok"},
			"odjrjznxpn",
			5,
			1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, closetTarget(tc.words, tc.target, tc.startIndex))
		})
	}
}

func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	res := math.MaxInt32
	for i := range words {
		if words[i] == target {
			targetIdx := i
			l := startIndex
			var ll int
			for l != targetIdx {
				l = (l + 1) % n
				ll++
			}
			var rr int
			r := startIndex
			for r != targetIdx {
				r = (r + (n - 1)) % n
				rr++
			}
			res = min(res, min(ll, rr))
		}
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
