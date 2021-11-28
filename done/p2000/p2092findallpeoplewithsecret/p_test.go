package p2092findallpeoplewithsecret

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAllPeople(t *testing.T) {
	for _, tc := range []struct {
		n           int
		meetings    [][]int
		firstPerson int
		want        []int
	}{
		{6, [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, 1, []int{0, 1, 2, 3, 5}},
		{4, [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, 3, []int{0, 1, 3}},
		{5, [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, 1, []int{0, 1, 2, 3, 4}},
		{6, [][]int{{0, 2, 1}, {1, 3, 1}, {4, 5, 1}}, 1, []int{0, 1, 2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findAllPeople(tc.n, tc.meetings, tc.firstPerson))
		})
	}
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	key := func(p1, p2 int32) [2]int32 {
		if p1 > p2 {
			p1, p2 = p2, p1
		}
		return [2]int32{p1, p2}
	}
	met := make(map[[2]int32]bool)

	h := make(MeetingTimeHeap, 0)
	h = append(h, meetingTime{int32(firstPerson), 0})
	met[key(0, int32(firstPerson))] = true
	adj := make([][]meetingTime, n)
	for _, meeting := range meetings {
		a, b, t := int32(meeting[0]), int32(meeting[1]), int32(meeting[2])
		adj[a] = append(adj[a], meetingTime{b, t})
		adj[b] = append(adj[b], meetingTime{a, t})
		if a == 0 {
			h = append(h, meetingTime{b, t})
		} else if b == 0 {
			h = append(h, meetingTime{a, t})
		}
	}
	heap.Init(&h)

	knows := make([]bool, n)
	knows[0] = true
	for len(h) > 0 {
		x := heap.Pop(&h).(meetingTime)
		knows[x.person] = true
		for _, nei := range adj[x.person] {
			if k := key(x.person, nei.person); nei.time >= x.time && !met[k] {
				heap.Push(&h, nei)
				met[k] = true
			}
		}
	}
	res := make([]int, 0)
	for i, doesKnow := range knows {
		if doesKnow {
			res = append(res, i)
		}
	}
	return res
}

type meetingTime struct {
	person int32
	time   int32
}

type MeetingTimeHeap []meetingTime

func (h MeetingTimeHeap) Len() int { return len(h) }
func (h MeetingTimeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MeetingTimeHeap) Less(i, j int) bool {
	return h[i].time < h[j].time
}
func (h *MeetingTimeHeap) Push(x interface{}) {
	*h = append(*h, x.(meetingTime))
}
func (h *MeetingTimeHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
