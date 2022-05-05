package p0155minstack

type MinStack struct {
	stack []item
	n     int
}

type item struct {
	val    int
	minVal int
}

func Constructor() MinStack {
	return MinStack{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(x int) {
	minVal := x
	if this.n == 0 {
		minVal = x
	} else {
		minVal = min(x, this.stack[this.n-1].minVal)
	}
	this.stack = append(this.stack, item{val: x, minVal: minVal})
	this.n++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.n-1]
	this.n--
}

func (this *MinStack) Top() int {
	return this.stack[this.n-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[this.n-1].minVal
}
