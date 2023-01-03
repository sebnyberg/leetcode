package p0710randompickwithblacklist

import "math/rand"

type Solution struct {
	mapping map[int]int
	m       int
	n       int
	x       int
	bl      map[int]bool
}

func Constructor(n int, blacklist []int) Solution {
	bl := make(map[int]bool)
	var m int
	for _, b := range blacklist {
		if b < n {
			m++
		}
		bl[b] = true
	}
	return Solution{
		n:       n,
		m:       m,
		x:       n - m,
		bl:      bl,
		mapping: make(map[int]int),
	}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.n - this.m)
	if !this.bl[x] {
		return x
	}
	if _, exists := this.mapping[x]; !exists {
		for this.bl[this.x] {
			this.x++
		}
		this.mapping[x] = this.x
		this.x++
	}
	return this.mapping[x]
}
