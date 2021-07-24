package p0362designhitcounter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

func Test_depthSum(t *testing.T) {
	nestedInt := []*NestedInteger{
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
	require.Equal(t, 8, depthSumInverse(nestedInt))
}

func depthSumInverse(nestedList []*NestedInteger) int {
	// Keep stack of lists
	ints := [][]*NestedInteger{nestedList}
	pos := []int{0}
	level := 0
	lens := []int{len(nestedList)}
	type levelNum struct {
		level int
		val   int
	}
	nums := make([]levelNum, 0, len(nestedList))
	for level >= 0 {
		if pos[level] == lens[level] {
			// pop
			ints = ints[:level]
			pos = pos[:level]
			lens = lens[:level]
			level--
			if level >= 0 {
				pos[level]++
			}
			continue
		}
		// Item is within current level
		cur := ints[level][pos[level]]
		if cur.IsInteger() {
			nums = append(nums, levelNum{level + 1, cur.GetInteger()})
			pos[level]++
		} else { // list => push
			l := cur.GetList()
			ints = append(ints, l)
			pos = append(pos, 0)
			lens = append(lens, len(l))
			level++
		}
	}
	// Find max depth
	maxDepth := 1
	for _, num := range nums {
		maxDepth = max(maxDepth, num.level)
	}
	// Inverse levels
	var sum int
	for i := range nums {
		sum += (maxDepth - nums[i].level + 1) * nums[i].val
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
