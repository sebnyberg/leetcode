package p0341flattenednestedlistiterator

import "testing"

func TestNestedIterator(t *testing.T) {
	i := []*NestedInteger{
		{
			children: []*NestedInteger{
				{val: 1},
				{val: 1},
			},
		},
		{val: 2},
		{
			children: []*NestedInteger{
				{val: 1},
				{val: 1},
			},
		},
	}
	it := Constructor(i)
	for it.HasNext() {
		it.Next()
	}

	i = []*NestedInteger{{children: []*NestedInteger{}}}
	it = Constructor(i)
	for it.HasNext() {
		it.Next()
	}
}

type NestedInteger struct {
	val      int
	children []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (i *NestedInteger) IsInteger() bool {
	return len(i.children) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (i *NestedInteger) GetInteger() int {
	return i.val
}

// Set this NestedInteger to hold a single integer.
func (i *NestedInteger) SetInteger(value int) {
	i.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (i *NestedInteger) Add(elem NestedInteger) {
	if len(i.children) == 0 {
		i.children = append(i.children, &NestedInteger{val: i.val})
	}
	i.children = append(i.children, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (i *NestedInteger) GetList() []*NestedInteger {
	return i.children
}

type NestedIterator struct {
	currentList []*NestedInteger
	parent      [][]*NestedInteger
	parentPos   []int
	pos         int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		currentList: nestedList,
		parent:      make([][]*NestedInteger, 0),
		parentPos:   make([]int, 0),
		pos:         0,
	}
}

func (this *NestedIterator) Next() int {
	val := this.currentList[this.pos].GetInteger()
	this.pos++
	return val
}

func (this *NestedIterator) dive() bool {
	if this.pos >= len(this.currentList) || this.currentList[this.pos].IsInteger() {
		return false
	}
	cur := this.currentList[this.pos].GetList()
	this.pos++
	this.parentPos = append(this.parentPos, this.pos)
	this.parent = append(this.parent, this.currentList)
	this.currentList = cur
	this.pos = 0
	return true
}

func (this *NestedIterator) bubble() bool {
	if this.pos < len(this.currentList) || len(this.parent) == 0 {
		return false
	}
	m := len(this.parent)
	this.currentList = this.parent[m-1]
	this.pos = this.parentPos[m-1]
	this.parent = this.parent[:m-1]
	this.parentPos = this.parentPos[:m-1]
	return true
}

func (this *NestedIterator) HasNext() bool {
	if this.pos >= len(this.currentList) && len(this.parent) == 0 {
		return false
	}
	for this.dive() || this.bubble() {
	}
	return this.pos < len(this.currentList)
}
