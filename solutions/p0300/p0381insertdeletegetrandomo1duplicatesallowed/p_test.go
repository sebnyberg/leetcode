package p0381insertdeletegetrandomo1duplicatesallowed

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRandomizedCollection(t *testing.T) {
	set := Constructor()
	insertRes := set.Insert(1)
	require.Equal(t, true, insertRes)
	insertRes = set.Insert(1)
	require.Equal(t, false, insertRes)
	insertRes = set.Insert(2)
	require.Equal(t, true, insertRes)
	count := make([]int, 3)
	for i := 0; i < 1000; i++ {
		rndRes := set.GetRandom()
		count[rndRes]++
	}
	require.InEpsilon(t, 0.6666, float64(count[1])/1000.0, 0.1)
	removeRes := set.Remove(1)
	require.Equal(t, true, removeRes)
	count = make([]int, 3)
	for i := 0; i < 1000; i++ {
		rndRes := set.GetRandom()
		count[rndRes]++
	}
	require.InEpsilon(t, 0.500, float64(count[1])/1000.0, 0.1)
}

type RandomizedCollection struct {
	itemIndices map[int]map[int]struct{}
	items       []int
	rnd         *rand.Rand
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		itemIndices: make(map[int]map[int]struct{}),
		items:       make([]int, 0),
		rnd:         rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Insert a value to the collection. Returns true if the collection did not
// already contain the specified element.
func (this *RandomizedCollection) Insert(val int) bool {
	_, exists := this.itemIndices[val]
	if !exists {
		this.itemIndices[val] = make(map[int]struct{}, 1)
	}
	idx := len(this.items)
	this.itemIndices[val][idx] = struct{}{}
	this.items = append(this.items, val)
	return !exists
}

// Removes a value from the collection. Returns true if the collection contained
// the specified element.
func (this *RandomizedCollection) Remove(val int) bool {
	indices, exists := this.itemIndices[val]
	if !exists {
		return false
	}

	lastIdx := len(this.items) - 1
	if _, exists := indices[lastIdx]; exists {
		delete(this.itemIndices[val], lastIdx)
		if len(this.itemIndices[val]) == 0 {
			delete(this.itemIndices, val)
		}
		this.items = this.items[:lastIdx]
		return true
	}

	// Pick first index for this value (guaranteed not to be last index)
	var idx int
	for k := range indices {
		idx = k
		break
	}

	// Delete index
	delete(this.itemIndices[val], idx)

	// Swap with last index
	this.items[idx], this.items[lastIdx] = this.items[lastIdx], this.items[idx]

	// Update position for the swapped element
	delete(this.itemIndices[this.items[idx]], lastIdx)
	this.itemIndices[this.items[idx]][idx] = struct{}{}

	// Delete val from list of item indices if this was the last existing element
	if len(this.itemIndices[val]) == 0 {
		delete(this.itemIndices, val)
	}
	this.items = this.items[:lastIdx]
	return true
}

// GetRandom returns a random element from the collection.
func (this *RandomizedCollection) GetRandom() int {
	return this.items[this.rnd.Intn(len(this.items))]
}
