package p0362designhitcounter

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHitCounter(t *testing.T) {
	c := Constructor()
	c.Hit(1)
	c.Hit(2)
	c.Hit(3)
	res := c.GetHits(4)
	require.Equal(t, 3, res)
	c.Hit(300)
	res = c.GetHits(300)
	require.Equal(t, 4, res)
	res = c.GetHits(301)
	require.Equal(t, 3, res)
}

type HitCounter struct {
	hits []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{}
}

// Hit records a hit at the provided timestamp. Timestamps are in monotonically
// increasing order.
func (this *HitCounter) Hit(timestamp int) {
	this.hits = append(this.hits, timestamp)
}

// GetHits returns the number of hits in the past 5 minutes from the provided
// timestamp.
func (this *HitCounter) GetHits(timestamp int) int {
	low := sort.SearchInts(this.hits, timestamp-299)
	high := sort.SearchInts(this.hits, timestamp+1)
	return high - low
}
