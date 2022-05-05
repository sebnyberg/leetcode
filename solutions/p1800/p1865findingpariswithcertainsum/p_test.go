package p1865findingpairswithcertainsum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindSumPairs(t *testing.T) {
	nums1 := []int{1, 1, 2, 2, 2, 3}
	nums2 := []int{1, 4, 5, 2, 5, 4}
	p := Constructor(nums1, nums2)
	res := p.Count(7)
	require.Equal(t, 8, res)
	p.Add(3, 2)
	res = p.Count(8)
	require.Equal(t, 2, res)
	res = p.Count(4)
	require.Equal(t, 1, res)
	p.Add(0, 1)
	p.Add(1, 1)
	res = p.Count(7)
	require.Equal(t, 11, res)
}

type intCount struct {
	val   int
	count int
}

type FindSumPairs struct {
	nums1Count []intCount
	nums2      []int
	nums2Count map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	var p FindSumPairs
	p.nums1Count = make([]intCount, 0)
	count := make(map[int]int)
	for _, n1 := range nums1 {
		count[n1]++
	}
	for val, c := range count {
		p.nums1Count = append(p.nums1Count, intCount{val, c})
	}
	sort.Slice(p.nums1Count, func(i, j int) bool {
		return p.nums1Count[i].val < p.nums1Count[j].val
	})
	p.nums2Count = make(map[int]int, len(nums2))
	for _, n2 := range nums2 {
		p.nums2Count[n2]++
	}
	p.nums2 = nums2
	return p
}

// Add val to nums2 at index index
func (this *FindSumPairs) Add(index int, val int) {
	this.nums2Count[this.nums2[index]]--
	this.nums2[index] += val
	this.nums2Count[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	var res int
	for _, it := range this.nums1Count {
		if it.val > tot {
			break
		}
		res += it.count * this.nums2Count[tot-it.val]
	}
	return res
}
