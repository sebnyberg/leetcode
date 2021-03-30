package p0354russiandollenvelopes

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxEnvelopes(t *testing.T) {
	for _, tc := range []struct {
		envelopes [][]int
		want      int
	}{
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.envelopes), func(t *testing.T) {
			require.Equal(t, tc.want, maxEnvelopes(tc.envelopes))
		})
	}
}

func maxEnvelopes(envelopes [][]int) int {
	// Sort envelopes by width and height (descending)
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, 0, len(envelopes))
	for _, envelope := range envelopes {
		i := sort.SearchInts(dp, envelope[1])
		if i == len(dp) {
			dp = append(dp, envelope[1])
		} else {
			dp[i] = envelope[1]
		}
	}
	return len(dp)
}
