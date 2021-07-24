package p0380insertdeletegetrandomo1

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRandomizedSet(t *testing.T) {
	set := Constructor()
	insertRes := set.Insert(1)
	require.Equal(t, true, insertRes)
	removeRes := set.Remove(2)
	require.Equal(t, false, removeRes)
	insertRes = set.Insert(2)
	require.Equal(t, true, insertRes)
	count := make([]int, 3)
	for i := 0; i < 1000; i++ {
		rndRes := set.GetRandom()
		count[rndRes]++
	}
	require.InEpsilon(t, 0.5, float64(count[1])/1000.0, 0.05)
	removeRes = set.Remove(1)
	require.Equal(t, true, removeRes)
	insertRes = set.Insert(2)
	require.Equal(t, false, insertRes)
	count = make([]int, 3)
	for i := 0; i < 1000; i++ {
		rndRes := set.GetRandom()
		count[rndRes]++
	}
	require.Equal(t, 1000, count[2])
}

type RandomizedSet struct {
	itemIndex map[int]int
	items     []int
	rnd       *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		itemIndex: make(map[int]int),
		items:     make([]int, 0),
		rnd:       rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Inserts a value to the set. Returns true if the set did not already contain
// the specified element.
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.itemIndex[val]; exists {
		return false
	}
	this.itemIndex[val] = len(this.items)
	this.items = append(this.items, val)
	return true
}

// Remove a value from the set. Returns true if the set contained the specified
// element.
func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.itemIndex[val]
	if !exists {
		return false
	}
	lastIdx := len(this.items) - 1
	this.items[idx], this.items[lastIdx] = this.items[lastIdx], this.items[idx]
	this.itemIndex[this.items[idx]] = idx
	this.items = this.items[:len(this.items)-1]
	delete(this.itemIndex, val)
	return true
}

// Get a random element from the set.
func (this *RandomizedSet) GetRandom() int {
	return this.items[this.rnd.Intn(len(this.items))]
}
