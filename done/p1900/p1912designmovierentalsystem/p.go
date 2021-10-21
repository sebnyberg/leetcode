package p1912designmovierentalsystem

import (
	"container/heap"
	"sort"
)

type MovieRentingSystem struct {
	rentMovies      *TopKHeap
	availableMovies map[uint16]*TopKHeap
	movies          map[location]*Movie
}

type location struct {
	shopID, movieID uint32
}

// const maxMovies = 10001

func Constructor(n int, entries [][]int) MovieRentingSystem {
	rs := MovieRentingSystem{
		rentMovies:      NewTopKHeap(5),
		availableMovies: make(map[uint16]*TopKHeap),
		movies:          make(map[location]*Movie),
	}
	for _, entry := range entries {
		shop, movie, price := entry[0], entry[1], entry[2]
		m := &Movie{
			id:    uint16(movie),
			price: uint16(price),
			shop:  uint32(shop),
		}
		rs.movies[location{uint32(shop), uint32(movie)}] = m
		if _, exists := rs.availableMovies[uint16(movie)]; !exists {
			rs.availableMovies[uint16(movie)] = NewTopKHeap(5)
		}
		rs.availableMovies[uint16(movie)].Push(m)
	}
	return rs
}

// Search finds the 5 cheapest shops that have an unrented copy of movie. Shops
// are sorted in ascending order, and in the case of the tie, the shop with the
// smaller index appears first.
func (this *MovieRentingSystem) Search(movie int) []int {
	res := make([]int, 0, 5)
	for _, m := range this.availableMovies[uint16(movie)].Top() {
		res = append(res, int(m.shop))
	}
	return res
}

// Rent rents an unrented copy of a movie from the specified shop.
func (this *MovieRentingSystem) Rent(shop int, movie int) {
	m := this.movies[location{uint32(shop), uint32(movie)}]
	this.availableMovies[uint16(movie)].Remove(m)
	this.rentMovies.Push(m)
}

// Drop drops off a previously rented copy of a movie at the specified shop.
func (this *MovieRentingSystem) Drop(shop int, movie int) {
	m := this.movies[location{uint32(shop), uint32(movie)}]
	this.rentMovies.Remove(m)
	this.availableMovies[uint16(movie)].Push(m)
}

// Report reports the cheapest 5 rented movies in a 2D list "res", where
// res[j] = [shop_j, movie_j] describes the jth cheapest rented movie was rented
// from the shop shop_j. The moviews in res should be sorted by price in
// ascending order, and in case of a tie, the smaller shop_j should appear
// first. If there is still a tie, the one with smaller movie_j should appear
// first.
func (this *MovieRentingSystem) Report() [][]int {
	res := make([][]int, 0)
	for _, m := range this.rentMovies.Top() {
		res = append(res, []int{int(m.shop), int(m.id)})
	}
	return res
}

type Movie struct {
	shop    uint32
	heapIdx uint32
	id      uint16
	price   uint16
	topK    bool
}

func (m *Movie) Less(other *Movie) bool {
	if m.price == other.price {
		if m.shop == other.shop {
			return m.id < other.id
		}
		return m.shop < other.shop
	}
	return m.price < other.price
}

func (m *Movie) Greater(other *Movie) bool {
	if m.price == other.price {
		if m.shop == other.shop {
			return m.id > other.id
		}
		return m.shop > other.shop
	}
	return m.price > other.price
}

type PriceHeap struct {
	items   []*Movie
	cmpFunc func(h []*Movie) func(i, j int) bool
}

func NewPriceHeap(cmpFunc func(h []*Movie) func(i, j int) bool) PriceHeap {
	h := PriceHeap{
		items:   make([]*Movie, 0),
		cmpFunc: cmpFunc,
	}
	return h
}

func (h PriceHeap) Len() int { return len(h.items) }
func (h PriceHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].heapIdx = uint32(i)
	h.items[j].heapIdx = uint32(j)
}
func (h PriceHeap) Less(i, j int) bool {
	return h.cmpFunc(h.items)(i, j)
}
func (h *PriceHeap) Push(x interface{}) {
	el := x.(*Movie)
	el.heapIdx = uint32(h.Len())
	h.items = append(h.items, el)
}

func (h *PriceHeap) Pop() interface{} {
	n := len(h.items)
	it := h.items[n-1]
	it.heapIdx = 0
	h.items = h.items[:n-1]
	return it
}
func (h *PriceHeap) Peek() *Movie {
	return h.items[0]
}

type TopKHeap struct {
	minHeap PriceHeap
	maxHeap PriceHeap
	k       int
}

func NewTopKHeap(k int) *TopKHeap {
	return &TopKHeap{
		k: k,
		minHeap: NewPriceHeap(func(h []*Movie) func(i, j int) bool {
			return func(i, j int) bool {
				return h[i].Less(h[j])
			}
		}),
		maxHeap: NewPriceHeap(func(h []*Movie) func(i, j int) bool {
			return func(i, j int) bool {
				return h[i].Greater(h[j])
			}
		}),
	}
}

func (h *TopKHeap) Push(m *Movie) {
	if h.maxHeap.Len() < h.k {
		heap.Push(&h.maxHeap, m)
		m.topK = true
		return
	}

	heap.Push(&h.minHeap, m)
	h.Balance()
}

func (h *TopKHeap) Balance() {
	for h.minHeap.Len() > 0 && h.maxHeap.Len() < h.k {
		// Move from min heap to max heap
		x := heap.Pop(&h.minHeap).(*Movie)
		x.topK = true
		heap.Push(&h.maxHeap, x)
	}

	for h.minHeap.Len() > 0 && h.minHeap.Peek().Less(h.maxHeap.Peek()) {
		// Move from min heap to max heap
		a := heap.Pop(&h.minHeap).(*Movie)
		a.topK = true
		b := heap.Pop(&h.maxHeap).(*Movie)
		b.topK = false
		heap.Push(&h.maxHeap, a)
		heap.Push(&h.minHeap, b)
	}

	for h.maxHeap.Len() > h.k {
		// Move from max heap to min heap
		x := heap.Pop(&h.maxHeap).(*Movie)
		x.topK = false
		heap.Push(&h.minHeap, x)
	}
}

// Remove removes the specified movie from the TopKHeap.
func (h *TopKHeap) Remove(m *Movie) {
	if m.topK {
		heap.Remove(&h.maxHeap, int(m.heapIdx))
	} else {
		heap.Remove(&h.minHeap, int(m.heapIdx))
	}
	m.topK = false
	h.Balance()
}

func (h *TopKHeap) Top() []*Movie {
	if h == nil {
		return []*Movie{}
	}
	res := make([]*Movie, h.maxHeap.Len())
	copy(res, h.maxHeap.items)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Less(res[j])
	})
	return res
}
