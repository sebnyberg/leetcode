package p1801numberofordersinthebacklog

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getNumberOfBacklogOrders(t *testing.T) {
	for _, tc := range []struct {
		orders [][]int
		want   int
	}{
		{[][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}, 6},
		{[][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}, 999999984},
	} {
		t.Run(fmt.Sprintf("%+v", tc.orders), func(t *testing.T) {
			require.Equal(t, tc.want, getNumberOfBacklogOrders(tc.orders))
		})
	}
}

const (
	buyType  = 0
	sellType = 1
)

const mod = 1000000007

func getNumberOfBacklogOrders(orders [][]int) int {
	buyMax := make(MaxHeap, 0)
	buyOrders := make(map[int]*orderPrice)

	sellMin := make(MinHeap, 0)
	sellOrders := make(map[int]*orderPrice)

	for _, ord := range orders {
		price, amount, orderType := ord[0], ord[1], ord[2]
		switch orderType {
		case buyType:
			for {
				if sellMin.Len() == 0 || amount == 0 || sellMin[0].price > price {
					break
				}
				if sellMin[0].count > amount {
					sellMin[0].count -= amount
					amount = 0
					break
				}
				amount -= sellMin[0].count
				delete(sellOrders, sellMin[0].price)
				heap.Pop(&sellMin)
			}
			if amount == 0 {
				continue
			} else if _, exists := buyOrders[price]; exists {
				buyOrders[price].count += amount
			} else {
				ord := &orderPrice{amount, price, 0}
				heap.Push(&buyMax, ord)
				buyOrders[price] = ord
			}
		case sellType:
			for {
				if buyMax.Len() == 0 || amount == 0 || buyMax[0].price < price {
					break
				}
				if buyMax[0].count > amount {
					buyMax[0].count -= amount
					amount = 0
					break
				}
				amount -= buyMax[0].count
				delete(buyOrders, buyMax[0].price)
				heap.Pop(&buyMax)
			}
			if amount == 0 {
				continue
			} else if _, exists := sellOrders[price]; exists {
				sellOrders[price].count += amount
			} else {
				ord := &orderPrice{amount, price, 0}
				heap.Push(&sellMin, ord)
				sellOrders[price] = ord
			}
		}
	}
	var totalCount int
	for _, o := range buyOrders {
		totalCount += o.count
		totalCount %= mod
	}
	for _, o := range sellOrders {
		totalCount += o.count
		totalCount %= mod
	}
	return totalCount
}

type orderPrice struct {
	count   int
	price   int
	heapIdx int
}

type MinHeap []*orderPrice

func (h *MinHeap) Len() int { return len(*h) }
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].price < (*h)[j].price
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h *MinHeap) Push(x interface{}) {
	n := len(*h)
	it := x.(*orderPrice)
	it.heapIdx = n
	*h = append(*h, it)
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

type MaxHeap []*orderPrice

func (h *MaxHeap) Len() int { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i].price > (*h)[j].price
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h *MaxHeap) Push(x interface{}) {
	n := len(*h)
	it := x.(*orderPrice)
	it.heapIdx = n
	*h = append(*h, it)
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}
