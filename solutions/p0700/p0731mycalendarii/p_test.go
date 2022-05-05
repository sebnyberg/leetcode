package p0731mycalendarii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = [][2]int{
	{12, 26},
	{70, 85},
	{55, 67},
	{2, 13},
	{3, 18},
	{91, 100},
	{13, 26},
	{17, 27},
	{41, 55},
	{15, 26},
	{50, 68},
	{34, 52},
	{95, 100},
	{23, 33},
	{89, 100},
	{27, 43},
	{80, 95},
	{97, 100},
	{28, 47},
	{45, 58},
	{76, 93},
	{56, 75},
	{91, 100},
	{61, 77},
	{36, 49},
	{18, 32},
	{96, 100},
	{96, 100},
	{67, 86},
	{46, 64},
	{95, 100},
	{17, 35},
	{8, 27},
	{4, 14},
	{30, 43},
	{74, 89},
	{77, 95},
	{98, 100},
	{31, 41},
	{35, 53}}

func TestCalendar(t *testing.T) {
	c := Constructor()
	// for _, x := range [][2]int{
	// 	{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10},
	// } {
	// 	c.Book(x[0], x[1])
	// }
	res := make([]bool, 0, 10)
	for _, x := range input[:10] {
		res = append(res, c.Book(x[0], x[1]))
	}
	res2 := c.Book(50, 68)
	require.Equal(t, true, res2)
}

type MyCalendarTwo struct {
	meetings [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	var doubleBookings [][2]int
	for _, interval := range this.meetings {
		l, r := interval[0], interval[1]
		if start >= r || end <= l {
			continue
		}
		doubleBookings = append(doubleBookings, interval)
	}
	// Check if double bookings overlap
	sort.Slice(doubleBookings, func(i, j int) bool {
		return doubleBookings[i][0] < doubleBookings[j][0]
	})
	for i := 0; i < len(doubleBookings)-1; i++ {
		if doubleBookings[i+1][0] < doubleBookings[i][1] {
			return false
		}
	}
	this.meetings = append(this.meetings, [2]int{start, end})
	return true
}
