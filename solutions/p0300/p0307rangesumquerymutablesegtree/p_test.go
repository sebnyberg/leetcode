package p0308rangesumquery2dmutablesegtree

type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	m := len(nums)
	n := 1
	for n < m {
		n <<= 1
	}
	var na NumArray
	na.n = n
	na.tree = make([]int, n*2)
	for i := range nums {
		na.tree[n+i] = nums[i]
	}
	for i := n - 1; i >= 1; i-- {
		na.tree[i] = na.tree[i*2] + na.tree[i*2+1]
	}
	return na
}

func (this *NumArray) Update(index int, val int) {
	delta := val - this.tree[this.n+index]
	for i := this.n + index; i >= 1; i /= 2 {
		this.tree[i] += delta
	}
}

func (this *NumArray) q(i, lo, hi, qlo, qhi int) int {
	if qhi <= lo || qlo >= hi {
		return 0
	}
	if qlo <= lo && qhi >= hi {
		return this.tree[i]
	}
	mid := lo + (hi-lo)/2
	return this.q(i*2, lo, mid, qlo, qhi) +
		this.q(i*2+1, mid, hi, qlo, qhi)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.q(1, 0, this.n, left, right+1)
}
