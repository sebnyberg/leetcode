package p0635designlogstoragesystem

import (
	"sort"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	// s := Constructor()
	// s.Put(1, "2017:01:01:23:59:59")
	// s.Put(2, "2017:01:01:22:59:59")
	// s.Put(3, "2016:01:01:00:00:00")
	// res := s.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year")
	// res = s.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour")

	s := Constructor()
	s.Put(1, "2017:01:01:23:59:59")
	s.Put(2, "2017:01:02:23:59:59")
	res := s.Retrieve("2017:01:01:23:59:58", "2017:01:02:23:59:58", "Second")
	_ = res
}

type Entry struct {
	ts time.Time
	id int
}

type LogSystem struct {
	logs []Entry
}

func (l *LogSystem) Less(i, j int) bool {
	return l.logs[i].ts.Before(l.logs[j].ts)
}

func (l *LogSystem) Len() int {
	return len(l.logs)
}

func (l *LogSystem) Swap(i, j int) {
	l.logs[i], l.logs[j] = l.logs[j], l.logs[i]
}

func Constructor() LogSystem {
	return LogSystem{
		logs: make([]Entry, 0, 100),
	}
}

func MustParse(s string) time.Time {
	ts, err := time.Parse("2006:01:02:15:04:05", s)
	if err != nil {
		panic(err)
	}
	return ts
}

const layout = "2006:01:03:15:04:05"

func (this *LogSystem) Put(id int, timestamp string) {
	this.logs = append(this.logs, Entry{
		ts: MustParse(timestamp),
		id: id,
	})
}

func (this *LogSystem) Retrieve(start string, end string, granularity string) []int {
	sort.Sort(this)

	startTS, endTS := MustParse(start), MustParse(end)
	switch granularity {
	case "Second":
		// endTS = endTS.Add(time.Second)
	case "Minute":
		startTS = time.Date(startTS.Year(), startTS.Month(), startTS.Day(), startTS.Hour(), startTS.Minute(), 0, 0, time.UTC)
		endTS = time.Date(endTS.Year(), endTS.Month(), endTS.Day(), endTS.Hour(), endTS.Minute(), 0, 0, time.UTC)
		endTS = endTS.Add(time.Minute)
	case "Hour":
		startTS = time.Date(startTS.Year(), startTS.Month(), startTS.Day(), startTS.Hour(), 0, 0, 0, time.UTC)
		endTS = time.Date(endTS.Year(), endTS.Month(), endTS.Day(), endTS.Hour(), 0, 0, 0, time.UTC)
		endTS = endTS.Add(time.Hour)
	case "Day":
		startTS = time.Date(startTS.Year(), startTS.Month(), startTS.Day(), 0, 0, 0, 0, time.UTC)
		endTS = time.Date(endTS.Year(), endTS.Month(), endTS.Day(), 0, 0, 0, 0, time.UTC)
		endTS = endTS.AddDate(0, 0, 1)
	case "Month":
		startTS = time.Date(startTS.Year(), startTS.Month(), 0, 0, 0, 0, 0, time.UTC)
		endTS = time.Date(endTS.Year(), endTS.Month(), 0, 0, 0, 0, 0, time.UTC)
		endTS = endTS.AddDate(0, 1, 0)
	case "Year":
		startTS = time.Date(startTS.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		endTS = time.Date(endTS.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		endTS = endTS.AddDate(1, 0, 0)
	}

	// fmt.Println(this.logs[0].ts)
	// fmt.Println(this.logs[1].ts)
	// fmt.Println(startTS)
	// fmt.Println(endTS)
	startIdx := sort.Search(len(this.logs), func(i int) bool {
		return this.logs[i].ts.After(startTS) || this.logs[i].ts.Equal(startTS)
	})
	endIdx := sort.Search(len(this.logs), func(i int) bool {
		return this.logs[i].ts.After(endTS)
	})
	res := make([]int, 0, endIdx-startIdx)
	for i := startIdx; i < endIdx; i++ {
		res = append(res, this.logs[i].id)
	}
	return res
}
