package p1235maximumprofitinjobscheduling

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jobScheduling(t *testing.T) {
	for _, tc := range []struct {
		startTime, endTime, profit []int
		want                       int
	}{
		{
			[]int{341, 22, 175, 424, 574, 687, 952, 439, 51, 562, 962, 890, 250, 47, 945, 914, 835, 937, 419, 343, 125, 809, 807, 959, 403, 861, 296, 39, 802, 562, 811, 991, 209, 375, 78, 685, 592, 409, 369, 478, 417, 162, 938, 298, 618, 745, 888, 463, 213, 351, 406, 840, 779, 299, 90, 846, 58, 235, 725, 676, 239, 256, 996, 362, 819, 622, 449, 880, 951, 314, 425, 127, 299, 326, 576, 743, 740, 604, 151, 391, 925, 605, 770, 253, 670, 507, 306, 294, 519, 184, 848, 586, 593, 909, 163, 129, 685, 481, 258, 764},
			[]int{462, 101, 820, 999, 900, 692, 991, 512, 655, 578, 996, 979, 425, 893, 975, 960, 930, 991, 987, 524, 208, 901, 841, 961, 878, 882, 412, 795, 937, 807, 957, 994, 963, 716, 608, 774, 681, 637, 635, 660, 750, 632, 948, 771, 943, 801, 985, 476, 532, 535, 929, 943, 837, 565, 375, 854, 174, 698, 820, 710, 566, 464, 997, 551, 884, 844, 830, 916, 970, 965, 585, 631, 785, 632, 892, 954, 803, 764, 283, 477, 970, 616, 794, 911, 771, 797, 776, 686, 895, 721, 917, 920, 975, 984, 996, 471, 770, 656, 977, 922},
			[]int{85, 95, 14, 72, 17, 3, 86, 65, 50, 50, 42, 75, 40, 87, 35, 78, 47, 74, 92, 10, 100, 29, 55, 57, 51, 34, 10, 96, 14, 71, 63, 99, 8, 37, 16, 71, 10, 71, 83, 88, 68, 79, 27, 87, 3, 58, 56, 43, 89, 31, 16, 9, 49, 84, 62, 30, 35, 7, 27, 34, 24, 33, 100, 25, 90, 79, 58, 21, 31, 30, 61, 46, 36, 45, 85, 62, 91, 54, 28, 63, 50, 69, 48, 36, 77, 39, 19, 97, 20, 39, 48, 72, 37, 67, 72, 46, 54, 37, 53, 30},
			998,
		},
		{[]int{4, 2, 4, 8, 2}, []int{5, 5, 5, 10, 8}, []int{1, 2, 8, 10, 4}, 18},
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
	type task struct {
		start, end, profit, maxProfit int
	}
	n := len(startTime)
	tasks := make([]task, n)
	for i := range startTime {
		tasks[i] = task{startTime[i], endTime[i], profit[i], profit[i]}
	}

	// Sort by end-time
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})
	tasks = append(tasks, task{0, 1e9 + 1, 0, 0}) // sentinel

	// Visit tasks one by one
	for i := 1; i < len(tasks)-1; i++ {
		// Find first task with an end time <= start time for current
		// This is done by binary searching for the first time greater
		gtIdx := sort.Search(i, func(j int) bool {
			return tasks[j].end > tasks[i].start
		})
		for j := gtIdx - 1; j >= 0 && tasks[j].end == tasks[gtIdx-1].end; j-- {
			tasks[i].maxProfit = max(tasks[i].maxProfit, tasks[i].profit+tasks[j].maxProfit)
		}
		tasks[i].maxProfit = max(tasks[i].maxProfit, tasks[i-1].maxProfit)
	}
	return tasks[len(tasks)-2].maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
