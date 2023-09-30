package p2747countzerorequestservers

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countServers(t *testing.T) {
	for i, tc := range []struct {
		n       int
		logs    [][]int
		x       int
		queries []int
		want    []int
	}{
		{
			3,
			leetcode.ParseMatrix("[[1,3],[2,6],[1,5]]"),
			5,
			[]int{10, 11},
			[]int{1, 2},
		},
		{
			6,
			leetcode.ParseMatrix("[[1,21]]"),
			10,
			[]int{24, 35, 28, 26, 20, 25, 16, 31, 12},
			[]int{5, 6, 5, 5, 6, 5, 6, 5, 6},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countServers(tc.n, tc.logs, tc.x, tc.queries))
		})
	}
}

func countServers(n int, logs [][]int, x int, queries []int) []int {
	// Sounds like we should add queries and logs to the same deque. When
	// popping a log, decrease the count of logs for that server. If the count
	// reaches zero, increment the "noLogsServerCount" by 1.
	type event struct {
		typ byte
		t   int
		val int
	}
	const typLog = 0
	const typQuery = 1
	timeline := make([]event, 0, len(logs)+len(queries))
	for _, l := range logs {
		timeline = append(timeline, event{typLog, l[1], l[0]})
	}
	for _, q := range queries {
		timeline = append(timeline, event{typQuery, q, 0})
	}

	sort.Slice(timeline, func(i, j int) bool {
		if timeline[i].t == timeline[j].t {
			return timeline[i].typ < timeline[j].typ
		}
		return timeline[i].t < timeline[j].t
	})
	deque := []event{}
	residx := make([]int, len(queries))
	for i := range residx {
		residx[i] = i
	}
	sort.Slice(residx, func(i, j int) bool {
		return queries[residx[i]] < queries[residx[j]]
	})

	var res []int
	serverCount := n
	serverLogs := make(map[int]int)
	for _, e := range timeline {
		for len(deque) > 0 && deque[0].t < e.t-x {
			a := deque[0]
			serverLogs[a.val]--
			if serverLogs[a.val] == 0 {
				serverCount++
			}
			deque = deque[1:]
		}
		if e.typ == typQuery {
			res = append(res, serverCount)
		} else {
			if serverLogs[e.val] == 0 {
				serverCount--
			}
			serverLogs[e.val]++
			deque = append(deque, e)
		}
	}
	res2 := make([]int, len(res))
	for i := range res {
		res2[residx[i]] = res[i]
	}
	return res2
}
