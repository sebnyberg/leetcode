package p0252meetingrooms

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canAttendMeetings(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      bool
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][]int{{7, 10}, {2, 4}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, canAttendMeetings(tc.intervals))
		})
	}
}

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t := intervals[0][1]
	for i, interval := range intervals {
		if i == 0 {
			continue
		}
		if interval[0] < t {
			return false
		}
		t = interval[1]
	}
	return true
}
