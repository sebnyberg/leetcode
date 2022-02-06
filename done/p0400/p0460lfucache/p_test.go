package p0460lfucache

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLFUCache(t *testing.T) {
	type any interface{}
	type action struct {
		name string
		inp  []any
		want []any
	}
	for _, tc := range []struct {
		name     string
		capacity int
		actions  []action
	}{
		{
			"simple4", 3,
			[]action{
				{"put", []any{2, 2}, []any{}},
				{"put", []any{1, 1}, []any{}},
				{"get", []any{2}, []any{2}},
				{"get", []any{1}, []any{1}},
				{"get", []any{2}, []any{2}},
				{"put", []any{3, 3}, []any{}},
				{"put", []any{4, 4}, []any{}},
				{"get", []any{3}, []any{-1}},
				{"get", []any{2}, []any{2}},
				{"get", []any{1}, []any{1}},
				{"get", []any{4}, []any{4}},
			},
		},
		{
			"simple3", 1,
			[]action{
				{"put", []any{2, 1}, []any{}},
				{"get", []any{2}, []any{1}},
				{"put", []any{3, 2}, []any{}},
				{"get", []any{2}, []any{-1}},
				{"get", []any{3}, []any{2}},
			},
		},
		{
			"zero capacity", 0,
			[]action{
				{"put", []any{0, 0}, []any{}},
				{"get", []any{0}, []any{-1}},
			},
		},
		{
			"simple", 2,
			[]action{
				{"put", []any{1, 1}, []any{}},
				{"put", []any{2, 2}, []any{}},
				{"get", []any{1}, []any{1}},
				{"put", []any{3, 3}, []any{}},
				{"get", []any{2}, []any{-1}},
				{"get", []any{3}, []any{3}},
				{"put", []any{4, 4}, []any{}},
				{"get", []any{1}, []any{-1}},
				{"get", []any{3}, []any{3}},
				{"get", []any{4}, []any{4}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cache := Constructor(tc.capacity)
			for i, act := range tc.actions {
				switch act.name {
				case "put":
					cache.Put(act.inp[0].(int), act.inp[1].(int))
				case "get":
					res := cache.Get(act.inp[0].(int))
					require.Equal(t, act.want[0].(int), res, "%v, %v", act.name, i)
				}
			}
		})
	}
}

type LFUCache struct {
	items     map[int]*cacheItem
	itemsHeap lfuHeap
	t         int
	capacity  int
}

type cacheItem struct {
	key        int
	val        int
	accessedAt int
	count      int
	heapIdx    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		items:     make(map[int]*cacheItem),
		itemsHeap: make(lfuHeap, 0),
		t:         0,
		capacity:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if _, exists := this.items[key]; !exists {
		return -1
	}
	item := this.items[key]
	item.accessedAt = this.t
	item.count++
	heap.Fix(&this.itemsHeap, item.heapIdx)
	this.t++
	return this.items[key].val
}

func (this *LFUCache) Put(key int, value int) {
	if _, exists := this.items[key]; exists {
		item := this.items[key]
		item.val = value
		item.accessedAt = this.t
		item.count++
		// Fix the heap based on the new value
		heap.Fix(&this.itemsHeap, item.heapIdx)
		return
	}

	if len(this.items) == this.capacity && len(this.items) > 0 {
		// Remove least frequently used, or least recently used remove
		remove := heap.Pop(&this.itemsHeap).(*cacheItem)
		delete(this.items, remove.key)
	}

	// Add item
	item := &cacheItem{
		key:        key,
		val:        value,
		accessedAt: this.t,
		count:      1,
		heapIdx:    0, // will be updated on push
	}
	heap.Push(&this.itemsHeap, item)
	this.items[key] = item
	this.t++
}

type lfuHeap []*cacheItem

func (h lfuHeap) Len() int { return len(h) }
func (h lfuHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h lfuHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].accessedAt < h[j].accessedAt
	}
	return h[i].count < h[j].count
}
func (h *lfuHeap) Push(x interface{}) {
	it := x.(*cacheItem)
	it.heapIdx = len(*h)
	*h = append(*h, it)
}
func (h *lfuHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
