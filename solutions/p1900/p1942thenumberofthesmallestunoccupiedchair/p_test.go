package p1942thenumberofthesmallestunoccupiedchair

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestChair(t *testing.T) {
	for _, tc := range []struct {
		times        [][]int
		targetFriend int
		want         int
	}{
		{[][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}},
			6, 2},
		{[][]int{{1, 4}, {2, 3}, {4, 6}}, 1, 1},
		{[][]int{{3, 10}, {1, 5}, {2, 6}}, 0, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.times), func(t *testing.T) {
			require.Equal(t, tc.want, smallestChair(tc.times, tc.targetFriend))
		})
	}
}

type arrival struct {
	id   int
	time int
}

type exit struct {
	id   int
	time int
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	arrivals := make([]arrival, n)
	exits := make([]exit, n)
	for i, t := range times {
		arrivals[i] = arrival{i, t[0]}
		exits[i] = exit{i, t[1]}
	}
	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i].time < arrivals[j].time
	})
	sort.Slice(exits, func(i, j int) bool {
		return exits[i].time < exits[j].time
	})
	var arrivalIdx, exitIdx int
	var chairsTaken [10001]bool
	friendsChair := make([]int, n)
	for i := range friendsChair {
		friendsChair[i] = -1
	}
	for {
		if arrivalIdx == n || exits[exitIdx].time <= arrivals[arrivalIdx].time {
			// Release taken chair
			chairsTaken[friendsChair[exits[exitIdx].id]] = false
			exitIdx++
		} else {
			arr := arrivals[arrivalIdx]
			// Take up new chair
			for i, taken := range chairsTaken {
				if !taken {
					if arr.id == targetFriend {
						return i
					}
					chairsTaken[i] = true
					friendsChair[arr.id] = i
					break
				}
			}
			arrivalIdx++
		}
	}
}
