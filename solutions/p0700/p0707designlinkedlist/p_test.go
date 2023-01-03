package p0707designlinkedlist

type listNode struct {
	next *listNode
	val  int
}

type MyLinkedList struct {
	n     int
	first *listNode
	last  *listNode
}

func Constructor() MyLinkedList {
	first := &listNode{val: -17}
	return MyLinkedList{
		first: first,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.n {
		return -1
	}
	prev := this.first
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	return prev.next.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.first.next = &listNode{
		next: this.first.next,
		val:  val,
	}
	if this.n == 0 {
		this.last = this.first.next
	}
	this.n++
}

func (this *MyLinkedList) AddAtTail(val int) {
	last := &listNode{
		val: val,
	}
	if this.n == 0 {
		this.first.next = last
	} else {
		this.last.next = last
	}
	this.last = last
	this.n++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.n {
		return
	}
	if index == this.n {
		this.AddAtTail(val)
		return
	}
	prev := this.first
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	if prev == nil {
		return
	}
	prev.next = &listNode{
		val:  val,
		next: prev.next,
	}
	this.n++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.n {
		return
	}
	prev := this.first
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	if index == this.n-1 {
		this.last = prev
	}
	this.n--
}
