package p0900rleiterator

import "testing"

func Test(t *testing.T) {
	i := RLEIterator{
		encoding: []int{3, 8, 0, 9, 2, 5},
	}
	var res []int
	res = append(res, i.Next(2))
	res = append(res, i.Next(1))
	res = append(res, i.Next(1))
	res = append(res, i.Next(2))
	_ = res
}

type RLEIterator struct {
	encoding []int
	i, j     int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{
		encoding: encoding,
	}
}

func (this *RLEIterator) Next(n int) int {
	res := -1
	for this.i < len(this.encoding) && this.j+n > this.encoding[this.i] {
		n -= this.encoding[this.i] - this.j
		if n == 0 && res == -1 {
			res = this.encoding[this.i+1]
		}
		this.i += 2
		this.j = 0
	}
	if this.i >= len(this.encoding) {
		return res
	}
	res = this.encoding[this.i+1]
	this.j += n
	return res
}
