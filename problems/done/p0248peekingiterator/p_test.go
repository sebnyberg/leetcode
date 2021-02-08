package p0248peekingiterator

// type Iterator struct {
// }

// func (it *Iterator) hasNext() bool {
// 	return true
// }

// func (it *Iterator) next() int {
// 	return 1
// }

// type PeekingIterator struct {
// 	iter       *Iterator
// 	peeked     bool
// 	peekedItem int
// }

// func Constructor(iter *Iterator) *PeekingIterator {
// 	return &PeekingIterator{
// 		iter: iter,
// 	}
// }

// func (this *PeekingIterator) hasNext() bool {
// 	if this.peeked {
// 		return true
// 	}
// 	return this.iter.hasNext()
// }

// func (this *PeekingIterator) next() int {
// 	if this.peeked {
// 		this.peeked = false
// 		return this.peekedItem
// 	}
// 	return this.iter.next()
// }

// func (this *PeekingIterator) peek() int {
// 	if this.peeked {
// 		return this.peekedItem
// 	}
// 	this.peeked = true
// 	this.peekedItem = this.iter.next()
// 	return this.peekedItem
// }
