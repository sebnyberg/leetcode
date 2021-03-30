package p0354russiandollenvelopes

import (
	"fmt"
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
	return 0
	// 	// Sort envelopes by width and height
	// 	byHeight := make([][]int, len(envelopes))
	// 	for i := range byHeight {
	// 		copy(byHeight[i], envelopes[i])
	// 	}
	// 	sort.Slice(byHeight, func(i, j int) bool { return byHeight[i][1] < byHeight[j][1] })
}
