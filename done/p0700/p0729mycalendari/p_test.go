package p0729mycalendari

type MyCalendar struct {
	meetings [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, interval := range this.meetings {
		l, r := interval[0], interval[1]
		if start >= r || end <= l {
			continue
		}
		return false
	}
	this.meetings = append(this.meetings, [2]int{start, end})
	return true
}
