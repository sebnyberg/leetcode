package p0519randomflipmatrix

import (
	"math/rand"
	"testing"
)

func TestSolution(t *testing.T) {
	s := Constructor(2, 2)
	res := s.Flip()
	res = s.Flip()
	res = s.Flip()
	s.Reset()
	res = s.Flip()
	_ = res
}

type Solution struct {
	m, n    int
	total   int
	flipped map[int]int
}

func Constructor(m int, n int) Solution {
	s := Solution{
		n:       n,
		m:       m,
		total:   m * n,
		flipped: make(map[int]int),
	}
	return s
}

func (this *Solution) Flip() []int {
	r := rand.Intn(this.total)
	this.total--
	x := r
	if v, exists := this.flipped[r]; exists {
		x = v
	}
	if v, exists := this.flipped[this.total]; exists {
		this.flipped[r] = v
	} else {
		this.flipped[r] = this.total
	}
	ii := x / this.n
	jj := x % this.n
	return []int{ii, jj}
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	for k := range this.flipped {
		delete(this.flipped, k)
	}
}
