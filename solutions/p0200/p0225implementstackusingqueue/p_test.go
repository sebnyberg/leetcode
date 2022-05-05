package p0225implementstackusingqueue

type MyStack struct {
	items []int
	n     int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		items: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
	this.n++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.items[this.n-1]
	this.items = this.items[:this.n-1]
	this.n--
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.items[this.n-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.n == 0
}
