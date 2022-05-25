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
	c.intervals = append(c.intervals, interval{-100, -99}, interval{1e9 + 1, 1e9 + 2})
	return c
}

func (this *CountIntervals) Add(left int, right int) {
	right++ // easier to work with open-ended intervals

	start := sort.Search(len(this.intervals), func(i int) bool {
		return left <= this.intervals[i].end
	})
	end := sort.Search(len(this.intervals), func(i int) bool {
		return this.intervals[i].start > right
	}) - 1

	// First scenario: <left> <right> [ )
	if right < this.intervals[start].start {
		// Insert on left side of current interval
		this.intervals = append(this.intervals, interval{})
		copy(this.intervals[start+1:], this.intervals[start:])
		this.intervals[start] = interval{left, right}
		this.count += right - left
		return
	} else if right > this.intervals[end].end {
		// [ ) <right>
		this.count += right - this.intervals[end].end
		this.intervals[end].end = right
	} else {
		// [ <right> )
	}

	if left < this.intervals[start].start {
		// <left> [ <right> ) or <left> [ ) <right>, etc
		this.count += this.intervals[start].start - left
		this.intervals[start].start = left
	} else {
		// <left> is inside the interval - no need to update
	}

	// Merge. These actions are a no-op when end == start
	for j := start + 1; j <= end; j++ {
		this.count += this.intervals[j].start - this.intervals[j-1].end
	}
	this.intervals[end].start = this.intervals[start].start
	copy(this.intervals[start:], this.intervals[end:])
	this.intervals = this.intervals[:len(this.intervals)-(end-start)]
}

func (this *CountIntervals) Count() int {
	return this.count
}
