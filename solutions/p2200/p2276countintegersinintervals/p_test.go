package p2276countintegersinintervals

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type actionType struct {
	name string
}

var actionTypeAdd = actionType{"add"}
var actionTypeCount = actionType{"count"}

type testAction struct {
	typ  actionType
	args []any
	want []any
}

func parseTest(s string) []testAction {
	rows := strings.Split(s, "\n")
	rows[0] = rows[0][1 : len(rows[0])-1]
	rows[1] = rows[1][2 : len(rows[1])-2]
	rows[2] = rows[2][1 : len(rows[2])-1]
	actionNames := strings.Split(rows[0], ",")
	actionArgs := strings.Split(rows[1], "],[")
	actionWant := strings.Split(rows[2], ",")
	res := []testAction{}
	for i, name := range actionNames {
		var a testAction
		if name[1:len(name)-1] == "count" {
			// Skip count actions for now to trigger panics.
			x, err := strconv.Atoi(actionWant[i])
			if err != nil {
				log.Fatalln(err)
			}
			a := testAction{
				actionTypeCount,
				nil,
				[]any{x},
			}
			res = append(res, a)
			continue
			// res = append(res, testAction{actionTypeAdd, nil, nil})
			// continue
		}
		for _, arg := range strings.Split(actionArgs[i], ",") {
			x, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalln(err)
			}
			a.args = append(a.args, x)
		}
		a.typ = actionTypeAdd
		res = append(res, a)
	}
	return res
}

func MustReadAll(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}

func TestCountIntervals(t *testing.T) {
	for i, tc := range [][]testAction{
		parseTest(MustReadAll("test2")),
		{
			{actionTypeCount, nil, []any{0}},
			{actionTypeAdd, []any{457, 717}, nil},
			{actionTypeAdd, []any{918, 927}, nil},
			{actionTypeCount, nil, []any{271}},
			{actionTypeAdd, []any{660, 675}, nil},
			{actionTypeCount, nil, []any{271}},
			{actionTypeCount, nil, []any{271}},
			{actionTypeAdd, []any{885, 905}, nil},
			{actionTypeCount, nil, []any{292}},
			{actionTypeCount, nil, []any{292}},
			{actionTypeAdd, []any{323, 416}, nil},
			{actionTypeAdd, []any{774, 808}, nil},
			{actionTypeCount, nil, []any{421}},
		},
		{
			{actionTypeCount, nil, []any{0}},
			{actionTypeAdd, []any{33, 49}, nil},
			{actionTypeAdd, []any{43, 47}, nil},
			{actionTypeCount, nil, []any{17}},
			{actionTypeCount, nil, []any{17}},
			{actionTypeAdd, []any{37, 37}, nil},
			{actionTypeAdd, []any{26, 38}, nil},
			{actionTypeAdd, []any{11, 11}, nil},
			{actionTypeCount, nil, []any{25}},
		},
		{
			{actionTypeAdd, []any{10, 27}, nil},
			{actionTypeAdd, []any{46, 50}, nil},
			{actionTypeAdd, []any{15, 35}, nil},
			{actionTypeAdd, []any{12, 32}, nil},
			{actionTypeAdd, []any{7, 15}, nil},
			{actionTypeAdd, []any{49, 49}, nil},
			{actionTypeCount, nil, []any{34}},
		},
		{
			{actionTypeCount, nil, []any{0}},
			{actionTypeAdd, []any{39, 44}, nil},
			{actionTypeAdd, []any{13, 49}, nil},
			{actionTypeCount, nil, []any{37}},
			{actionTypeCount, nil, []any{37}},
			{actionTypeAdd, []any{47, 50}, nil},
		},
		{
			{actionTypeAdd, []any{2, 3}, nil},
			{actionTypeAdd, []any{7, 10}, nil},
			{actionTypeCount, nil, []any{6}},
			{actionTypeAdd, []any{5, 8}, nil},
			{actionTypeCount, nil, []any{8}},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := Constructor()
			for j, a := range tc {
				switch a.typ {
				case actionTypeAdd:
					c.Add(a.args[0].(int), a.args[1].(int))
				case actionTypeCount:
					if a.want[0].(int) != c.Count() {
						t.Log(j)
						t.FailNow()
					}
				}
			}
		})
	}
}

type CountIntervals struct {
	intervals []interval
	count     int
}

type interval struct {
	start, end int
}

func Constructor() CountIntervals {
	c := CountIntervals{
		intervals: make([]interval, 0, 100),
	}
	// Add sentinel values to make indexing easier
	c.intervals = append(c.intervals,
		interval{-100, -99},
		interval{1e9 + 1, 1e9 + 2},
	)
	return c
}

func (this *CountIntervals) Add(left int, right int) {
	right++ // easier to work with open-ended intervals

	// Binary search for first and second interval.
	// These may be the same interval
	first := sort.Search(len(this.intervals), func(i int) bool {
		return left <= this.intervals[i].end
	})
	second := sort.Search(len(this.intervals), func(i int) bool {
		return this.intervals[i].start > right
	}) - 1

	// Cases to consider:
	//
	// <right> <left> [ )
	// <left> [ <right> )
	// <left> [ ) <right>
	// [ <left> <right> )
	// [ <left> ) <right>

	if right < this.intervals[first].start {
		// (1) <left> <right> [ )
		// Create new interval and return
		this.intervals = append(this.intervals, interval{})
		copy(this.intervals[first+1:], this.intervals[first:])
		this.intervals[first] = interval{left, right}
		this.count += right - left
		return
	} else if right > this.intervals[second].end {
		// Right falls outside the interval - update right boundary
		// E.g.
		// (3b) and (5b)
		// Adjust right boundary of the second interval
		this.count += right - this.intervals[second].end
		this.intervals[second].end = right
	} else {
		// (2a) (2b) (4a) (4b)
		// right is inside the interval - no need to update
		// [ <right> )
	}

	if left < this.intervals[first].start {
		// Example cases:
		// <left> [ <right> ) or <left> [ ) <right>, etc
		this.count += this.intervals[first].start - left
		this.intervals[first].start = left
	} else {
		// <left> is inside the interval - no need to update
	}

	// If first != second, merge intervals. These actions will do nothing when
	// start == end so no need to check.
	for j := first + 1; j <= second; j++ {
		this.count += this.intervals[j].start - this.intervals[j-1].end
	}
	this.intervals[second].start = this.intervals[first].start
	copy(this.intervals[first:], this.intervals[second:])
	this.intervals = this.intervals[:len(this.intervals)-(second-first)]
}

func (this *CountIntervals) Count() int {
	return this.count
}
