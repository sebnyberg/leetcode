package p0641designcircularqueue

import "testing"

func TestA(t *testing.T) {
	a := Constructor(77)
	a.InsertFront(89)
	a.GetRear()
	a.DeleteLast()
	a.GetRear()
	a.InsertFront(19)
	a.InsertFront(23)
	a.InsertFront(23)
	a.InsertFront(82)
	a.IsFull()
}

type node struct {
	prev *node
	next *node
	val  int
}

type MyCircularDeque struct {
	front *node
	last  *node
	n     int
	k     int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		k: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.n >= this.k {
		return false
	}
	front := &node{val: value}
	if this.n == 0 {
		this.last = front
	} else {
		front.next = this.front
		this.front.prev = front
	}
	this.front = front
	this.n++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.n >= this.k {
		return false
	}
	last := &node{val: value}
	if this.n == 0 {
		this.front = last
	} else {
		last.prev = this.last
		this.last.next = last
	}
	this.last = last
	this.n++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.n == 0 {
		return false
	}
	if this.n >= 2 {
		this.front = this.front.next
		this.front.prev = nil
	}
	this.n--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.n == 0 {
		return false
	}
	if this.n >= 2 {
		this.last = this.last.prev
		this.last.next = nil
	}
	this.n--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.n == 0 {
		return -1
	}
	return this.front.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.n == 0 {
		return -1
	}
	return this.last.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.n >= this.k
}
