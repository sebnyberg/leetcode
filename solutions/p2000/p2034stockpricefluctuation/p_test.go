package p2034stockpricefluctuation

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStockPrice(t *testing.T) {
	// sp := Constructor()
	// sp.Update(1, 10)
	// sp.Update(2, 5)
	// require.Equal(t, 5, sp.Current())
	// require.Equal(t, 10, sp.Maximum())
	// sp.Update(1, 3)
	// require.Equal(t, 5, sp.Maximum())
	// sp.Update(4, 2)
	// require.Equal(t, 2, sp.Minimum())

	sp := Constructor()
	sp.Update(1, 100)
	sp.Update(10, 10)
	sp.Update(1, 1)
	require.Equal(t, 1, sp.Minimum())
}

type StockPrice struct {
	minPrice   MinHeap
	maxPrice   MaxHeap
	stockForTs map[int]*StockQuote
	maxTs      int
}

func Constructor() StockPrice {
	return StockPrice{
		minPrice:   make(MinHeap, 0),
		maxPrice:   make(MaxHeap, 0),
		stockForTs: make(map[int]*StockQuote),
		maxTs:      -1,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.maxTs = max(this.maxTs, timestamp)
	if ticker, exists := this.stockForTs[timestamp]; exists {
		ticker.price = price
		heap.Fix(&this.maxPrice, ticker.maxIdx)
		heap.Fix(&this.minPrice, ticker.minIdx)
		return
	}
	this.stockForTs[timestamp] = &StockQuote{
		price: price,
	}
	heap.Push(&this.maxPrice, this.stockForTs[timestamp])
	heap.Push(&this.minPrice, this.stockForTs[timestamp])
}

func (this *StockPrice) Current() int {
	return this.stockForTs[this.maxTs].price
}

func (this *StockPrice) Maximum() int {
	return this.maxPrice[0].price
}

func (this *StockPrice) Minimum() int {
	return this.minPrice[0].price
}

type StockQuote struct {
	minIdx int
	maxIdx int
	price  int
}

type MinHeap []*StockQuote

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].minIdx = i
	h[j].minIdx = j
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].price < h[j].price
}
func (h *MinHeap) Push(x interface{}) {
	sq := x.(*StockQuote)
	sq.minIdx = len(*h)
	*h = append(*h, sq)
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

type MaxHeap []*StockQuote

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].maxIdx = i
	h[j].maxIdx = j
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].price > h[j].price
}
func (h *MaxHeap) Push(x interface{}) {
	sq := x.(*StockQuote)
	sq.maxIdx = len(*h)
	*h = append(*h, sq)
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
