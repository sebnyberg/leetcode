package p0495teemoattacking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPoisonedDuration(t *testing.T) {
	for _, tc := range []struct {
		timeSeries []int
		duration   int
		want       int
	}{
		{[]int{1, 4}, 2, 4},
		{[]int{1, 2}, 2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.timeSeries), func(t *testing.T) {
			require.Equal(t, tc.want, findPoisonedDuration(tc.timeSeries, tc.duration))
		})
	}
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	var total int
	for i, t := range timeSeries {
		if i < len(timeSeries)-1 {
			total += min(timeSeries[i+1], t+duration) - t
		} else {
			total += duration
		}
	}
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
