package p1010pairofsongswithtotaldurationsdivisibleby60

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numPairsDivisibleBy60(t *testing.T) {
	for _, tc := range []struct {
		time []int
		want int
	}{
		{[]int{30, 20, 150, 100, 40}, 3},
		{[]int{60, 60, 60}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.time), func(t *testing.T) {
			require.Equal(t, tc.want, numPairsDivisibleBy60(tc.time))
		})
	}
}

func numPairsDivisibleBy60(time []int) int {
	var timeCount [60]int
	var result int
	for _, t := range time {
		result += timeCount[(60-(t%60))%60]
		timeCount[t%60]++
	}
	return result
}
