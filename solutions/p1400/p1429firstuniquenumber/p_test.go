package p1427firstuniquenumber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstUnique(t *testing.T) {
	fu := Constructor([]int{1, 1, 2, 3})
	require.Equal(t, 2, fu.ShowFirstUnique())
	fu.Add(1)
	require.Equal(t, 2, fu.ShowFirstUnique())
	fu.Add(3)
	fu.Add(3)
	require.Equal(t, 2, fu.ShowFirstUnique())
	fu.Add(2)
	require.Equal(t, -1, fu.ShowFirstUnique())
}

type FirstUnique struct {
	uniqueNums []int
	numToIdx   map[int]int
	start      int
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		uniqueNums: nums,
		numToIdx:   make(map[int]int),
	}
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}
	var insertPos int
	for _, num := range nums {
		if numCount[num] == 1 {
			nums[insertPos] = num
			fu.numToIdx[num] = insertPos
			insertPos++
		} else {
			fu.numToIdx[num] = -1
		}
	}
	fu.uniqueNums = fu.uniqueNums[:insertPos]
	return fu
}

func (this *FirstUnique) prune() {
	for this.start != len(this.uniqueNums) && this.uniqueNums[this.start] == -1 {
		this.start++
	}
}

func (this *FirstUnique) ShowFirstUnique() int {
	this.prune()
	if this.start == len(this.uniqueNums) {
		return -1
	}
	return this.uniqueNums[this.start]
}

func (this *FirstUnique) Add(value int) {
	if idx, exists := this.numToIdx[value]; !exists {
		this.numToIdx[value] = len(this.uniqueNums)
		this.uniqueNums = append(this.uniqueNums, value)
	} else if idx > -1 {
		this.uniqueNums[idx] = -1
	}
}
