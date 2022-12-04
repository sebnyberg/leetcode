package p1024videostitching

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_videoStitching(t *testing.T) {
	for i, tc := range []struct {
		clips [][]int
		time  int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[8,10],[17,39],[18,19],[8,16],[13,35],[33,39],[11,19],[18,35]]"),
			20, 0,
		},
		{
			leetcode.ParseMatrix("[[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]"),
			10, 3,
		},
		{
			leetcode.ParseMatrix("[[0,1],[1,2]]"),
			5, -1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, videoStitching(tc.clips, tc.time))
		})
	}
}

func videoStitching(clips [][]int, time int) int {
	// When there are multiple options, choose the longest clip
	// When there is no option, give up
	//
	// Sort clips by start-times.
	// While clips start <= the current timestamp, find the max end time for any
	// clip. If the max end time does not include the current timestamp, exit.
	//
	var res int
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] <= clips[j][0]
	})
	var j int
	var t int
	for t < time {
		bestEndTime := -1
		for j < len(clips) && clips[j][0] <= t {
			bestEndTime = max(bestEndTime, clips[j][1])
			j++
		}
		if bestEndTime < t {
			return -1
		}
		t = bestEndTime
		res++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
