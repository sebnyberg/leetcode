package p0729mycalendar1

import "math"

type MyCalendar struct {
	bookings IntervalTree
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.bookings.IntersectsInterval(start, end) {
		return false
	}
	this.bookings.Insert(start, end)
	return true
}

// IntervalTree is a modified AVL-tree which can be used to perform O(nlogn)
// lookups of intervals.
type IntervalTree struct {
	root *intervalNode
}

// Insert the interval [min, max) i.e. lower-inclusive, upper-exclusive.
func (t *IntervalTree) Insert(start, end int) {
	if t.root == nil {
		t.root = &intervalNode{
			start:   start,
			end:     end,
			treeMax: end,
			height:  1,
		}
		return
	}
	t.root = t.root.insert(start, end)
}

func (t *IntervalTree) FindInterval(start, end int) {
	t.root.findInterval(start, end)
}

// IntersectsInterval checks if the tree contains an interval which intersects
// the provided interval [start,end).
func (t *IntervalTree) IntersectsInterval(start, end int) bool {
	return t.root.intersectsInterval(start, end)
}

type intervalNode struct {
	start, end int
	left       *intervalNode
	right      *intervalNode
	height     int // Height of tree
	treeMax    int // Tree maximum
}

// insert inserts the provided value and updates the height of this tree.
// Calling this function may cause a rebalance operation, in which case
// the returned value is the new root of the sub-tree. Ensure that the caller
// updates references accordingly.
func (n *intervalNode) insert(lower, upper int) *intervalNode {
	if lower < n.start {
		if n.left == nil {
			n.left = &intervalNode{
				start:   lower,
				end:     upper,
				treeMax: upper,
				height:  1,
			}
		} else {
			n.left = n.left.insert(lower, upper)
		}
	} else {
		if n.right == nil {
			n.right = &intervalNode{
				start:   lower,
				end:     upper,
				treeMax: upper,
				height:  1,
			}
		} else {
			n.right = n.right.insert(lower, upper)
		}
	}
	n.treeMax = max(n.treeMax, upper)
	n.updateHeight()

	balance := n.right.getHeight() - n.left.getHeight()

	if balance < -1 {
		// LL
		if n.left.left.getHeight() > n.left.right.getHeight() {
			return n.rotateRight()
		} else { // LR
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	} else if balance > 1 {
		// RR
		if n.right.right.getHeight() > n.right.left.getHeight() {
			return n.rotateLeft()
		} else { // RL
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}
	return n
}

func (n *intervalNode) updateHeight() {
	if n != nil {
		n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
	}
}

func (t *intervalNode) getHeight() int {
	if t == nil {
		return 0
	}
	return t.height
}

func (t *intervalNode) getTreeMax() int {
	if t == nil {
		return math.MinInt32
	}
	return t.treeMax
}

func (n *intervalNode) rotateLeft() *intervalNode {
	newRoot := n.right
	n.right.left, n.right = n, n.right.left
	newRoot.updateHeight()
	n.updateHeight()
	newRoot.treeMax = n.treeMax
	n.treeMax = max(n.end, max(n.left.getTreeMax(), n.right.getTreeMax()))
	return newRoot
}

func (n *intervalNode) rotateRight() *intervalNode {
	newRoot := n.left
	n.left.right, n.left = n, n.left.right
	newRoot.updateHeight()
	n.updateHeight()
	newRoot.treeMax = n.treeMax
	return newRoot
}

func (n *intervalNode) findInterval(start, end int) *intervalNode {
	if n == nil {
		return nil
	}
	if n.start == start && n.end == end {
		return n
	}
	if n.left.findInterval(start, end) != nil {
		return n.left
	}
	if n.right.findInterval(start, end) != nil {
		return n.right
	}
	return nil
}

func (n *intervalNode) intersectsInterval(start, end int) bool {
	if n == nil {
		return false
	}
	if start < n.end && end > n.start {
		return true
	}
	if n.left != nil && n.left.getTreeMax() > start {
		return n.left.intersectsInterval(start, end)
	}
	return n.right.intersectsInterval(start, end)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
