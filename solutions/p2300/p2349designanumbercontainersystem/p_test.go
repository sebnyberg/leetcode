package p2349designanumbercontainersystem

import (
	"container/heap"
)

type entry struct {
	idx     int
	num     int
	heapIdx int
}

type NumberContainers struct {
	entries map[int]*entry
	nums    map[int]*entryHeap
}

func Constructor() NumberContainers {
	c := NumberContainers{
		entries: make(map[int]*entry),
		nums:    make(map[int]*entryHeap),
	}
	return c
}

func (this *NumberContainers) Change(index int, number int) {
	if _, exists := this.nums[number]; !exists {
		this.nums[number] = &entryHeap{}
	}
	if e, exists := this.entries[index]; exists {
		heap.Remove(this.nums[e.num], e.heapIdx)
		n := this.nums[number].Len()
		this.entries[index].heapIdx = n
		this.entries[index].num = number
	} else {
		n := this.nums[number].Len()
		this.entries[index] = &entry{
			idx:     index,
			num:     number,
			heapIdx: n,
		}
	}
	heap.Push(this.nums[number], this.entries[index])
}

func (this *NumberContainers) Find(number int) int {
	if _, exists := this.nums[number]; !exists || this.nums[number].Len() == 0 {
		return -1
	}
	return (*this.nums[number])[0].idx
}

type entryHeap []*entry

func (h entryHeap) Len() int { return len(h) }
func (h entryHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h entryHeap) Less(i, j int) bool {
	if h[i].num == h[j].num {
		return h[i].idx < h[j].idx
	}
	return h[i].num < h[j].num
}
func (h *entryHeap) Push(x interface{}) {
	*h = append(*h, x.(*entry))
}

func (h *entryHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
