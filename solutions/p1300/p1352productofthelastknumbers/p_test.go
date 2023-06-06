package p1352productofthelastknumbers

type ProductOfNumbers struct {
	res  []int
	zero int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		res:  []int{0},
		zero: 0,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.zero = len(this.res)
	}
	if len(this.res) == 0 || this.res[len(this.res)-1] == 0 {
		this.res = append(this.res, num)
	} else {
		this.res = append(this.res, num*this.res[len(this.res)-1])
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.res)
	if this.zero >= n-k {
		return 0
	}
	div := max(1, this.res[n-k-1])
	return this.res[n-1] / div
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
