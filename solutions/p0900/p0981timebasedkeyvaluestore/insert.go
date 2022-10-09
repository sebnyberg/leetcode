package p0981timebasedkeyvaluestore

import (
	"sort"
)

type TimeMap struct {
	timestamps map[string][]int
	values     map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		timestamps: make(map[string][]int),
		values:     make(map[string][]string),
	}
}

func (this *TimeMap) Set(key string, v string, timestamp int) {
	this.timestamps[key] = append(this.timestamps[key], timestamp)
	this.values[key] = append(this.values[key], v)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if this.timestamps[key] == nil {
		return ""
	}
	tss := this.timestamps[key]
	i := sort.Search(len(tss), func(i int) bool {
		return tss[i] > timestamp
	})
	if i == 0 {
		return ""
	}
	return this.values[key][i-1]
}
