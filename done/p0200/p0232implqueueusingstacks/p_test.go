package p0232implqueueusingstacks

type MyQueue struct {
	items []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		items: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyQueue) Push(x int) {
	this.items = append(this.items, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.items[0]
	this.items = this.items[1:]
	return val
}

/** Get the top element. */
func (this *MyQueue) Peek() int {
	return this.items[0]
}

/** Returns whether the stack is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.items) == 0
}
