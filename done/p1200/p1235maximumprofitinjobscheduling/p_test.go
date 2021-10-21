package p1235maximumprofitinjobscheduling

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jobScheduling(t *testing.T) {
	for _, tc := range []struct {
		startTime, endTime, profit []int
		want                       int
	}{
		{[]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}, 120},
		{[]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}, 150},
		{[]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startTime), func(t *testing.T) {
			require.Equal(t, tc.want, jobScheduling(tc.startTime, tc.endTime, tc.profit))
		})
	}
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {

}
