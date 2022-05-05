package p1817findingtheusersactiveminutes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findingUsersActiveMinutes(t *testing.T) {
	for _, tc := range []struct {
		logs [][]int
		k    int
		want []int
	}{
		{[][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5, []int{0, 2, 0, 0, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.logs), func(t *testing.T) {
			require.Equal(t, tc.want, findingUsersActiveMinutes(tc.logs, tc.k))
		})
	}
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userActiveTimes := make(map[int]map[int]struct{})
	for _, log := range logs {
		id, time := log[0], log[1]
		if _, exists := userActiveTimes[id]; !exists {
			userActiveTimes[id] = make(map[int]struct{})
		}
		userActiveTimes[id][time] = struct{}{}
	}

	res := make([]int, k)
	for _, activeTimes := range userActiveTimes {
		res[len(activeTimes)-1]++
	}
	return res
}
