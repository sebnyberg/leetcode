package p0253meetingrooms2

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMeetingRooms(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{15, 16}, {10, 15}, {16, 25}}, 1},
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][]int{{7, 10}, {2, 4}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, minMeetingRooms(tc.intervals))
		})
	}
}

type meetingEventType byte

const (
	meetingEventTypeStarted = 0
	meetingEventTypeEnded   = 1
)

type MeetingEvent struct {
	time      int
	eventType meetingEventType
}

type MeetingEvents []MeetingEvent

func (e *MeetingEvents) Push(x interface{}) {
	(*e) = append((*e), x.(MeetingEvent))
}

func (e *MeetingEvents) Pop() interface{} {
	n := len(*e)
	el := (*e)[n-1]
	(*e) = (*e)[:n-1]
	return el
}

func (e *MeetingEvents) Less(i, j int) bool {
	if (*e)[i].time == (*e)[j].time {
		return (*e)[i].eventType == meetingEventTypeEnded
	}
	return (*e)[i].time < (*e)[j].time
}
func (e MeetingEvents) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *MeetingEvents) Len() int { return len(*e) }

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	events := make(MeetingEvents, n*2)
	for i := range intervals {
		events[i] = MeetingEvent{intervals[i][0], meetingEventTypeStarted}
		events[n+i] = MeetingEvent{intervals[i][1], meetingEventTypeEnded}
	}
	heap.Init(&events)
	var maxRooms int
	var curRooms int
	for len(events) > 0 {
		e := heap.Pop(&events).(MeetingEvent)
		if e.eventType == meetingEventTypeStarted {
			curRooms++
		} else {
			curRooms--
		}
		maxRooms = max(maxRooms, curRooms)
	}

	return maxRooms
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
