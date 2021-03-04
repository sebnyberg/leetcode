package p0146lrucache

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"
)

type action int

const (
	put action = iota
	get
)

func Test_LRUCache(t *testing.T) {
	c := Constructor(2)
	actions := []struct {
		a    action
		args []int
		want int
	}{
		{put, []int{1, 1}, 0},
		{put, []int{2, 2}, 0},
		{get, []int{1}, 1},
		{put, []int{3, 3}, 0},
		{get, []int{2}, -1},
		{put, []int{4, 4}, 0},
		{get, []int{1}, -1},
		{get, []int{3}, 3},
		{get, []int{4}, 4},
	}
	for _, a := range actions {
		switch a.a {
		case put:
			c.Put(a.args[0], a.args[1])
		case get:
			require.Equal(t, a.want, c.Get(a.args[0]))
		}
	}
}

type ListNode struct {
	Next *ListNode
	Prev *ListNode
	Val  int
}

type LRUCache struct {
	capacity  int
	items     map[int]*list.Element
	evictList *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		items:     make(map[int]*list.Element),
		evictList: list.New(),
	}
}

type entry struct {
	key   int
	value int
}

func (this *LRUCache) Get(key int) int {
	if el, exists := this.items[key]; exists {
		this.evictList.MoveToFront(el)
		return el.Value.(entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if el, exists := this.items[key]; exists {
		if val, ok := this.items[key].Value.(entry); ok {
			val.value = value
			this.items[key].Value = val
		}
		this.evictList.MoveToFront(el)
		return
	}
	if this.capacity == this.evictList.Len() {
		last := this.evictList.Back()
		delete(this.items, last.Value.(entry).key)
		this.evictList.Remove(last)
	}
	this.items[key] = this.evictList.PushFront(entry{key: key, value: value})
}
