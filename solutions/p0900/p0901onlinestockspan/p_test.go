package p0901onlinestockspan

import (
	"math"
)

type StockSpanner struct {
	prices []int
	count  []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: []int{math.MaxInt32},
		count:  []int{0},
	}
}

func (this *StockSpanner) Next(price int) int {
	cur := 1
	for this.prices[len(this.prices)-1] <= price {
		cur += this.count[len(this.count)-1]
		this.prices = this.prices[:len(this.prices)-1]
		this.count = this.count[:len(this.count)-1]
	}
	this.prices = append(this.prices, price)
	this.count = append(this.count, cur)
	return cur
}
