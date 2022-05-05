package p0384shuffleanarray

import (
	"math/rand"
	"testing"
)

func TestShuffler(t *testing.T) {
	in := []int{1, 2, 3}
	s := Constructor(in)
	r := s.Shuffle()
	_ = r
}

type Solution struct {
	nums     []int
	shuffled []int
	n        int
}

func Constructor(nums []int) Solution {
	shuffled := make([]int, len(nums))
	copy(shuffled, nums)
	return Solution{
		nums:     nums,
		shuffled: shuffled,
		n:        len(nums),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := this.n - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		this.shuffled[i], this.shuffled[j] = this.shuffled[j], this.shuffled[i]
	}
	return this.shuffled
}
