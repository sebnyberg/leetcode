package p0732mycalendarii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	c := Constructor()
	require.Equal(t, 1, c.Book(10, 20))
	require.Equal(t, 1, c.Book(50, 60))
	require.Equal(t, 2, c.Book(10, 40))
	require.Equal(t, 3, c.Book(5, 15))
	require.Equal(t, 3, c.Book(5, 10))
	require.Equal(t, 3, c.Book(25, 55))
}

type MyCalendarThree struct {
	// times contain negative entries for ends and positive entries for starts
	times []int
}

func Constructor() MyCalendarThree {
	var c MyCalendarThree
	c.times = make([]int, 0, 3)
	c.times = append(c.times, 1e9+7, -(1e9 + 9))
	return c
}

func (this *MyCalendarThree) Book(start int, end int) int {
	// Find ordered insert position start and end
	i := sort.Search(len(this.times), func(i int) bool {
		if abs(this.times[i]) == start {
			return this.times[i] > 0
		}
		return abs(this.times[i]) >= start
	})
	this.times = append(this.times, 0)
	copy(this.times[i+1:], this.times[i:])
	this.times[i] = start

	j := sort.Search(len(this.times), func(i int) bool {
		return abs(this.times[i]) >= end
	})
	this.times = append(this.times, 0)
	copy(this.times[j+1:], this.times[j:])
	this.times[j] = -end

	var k int
	var res int
	for m := 0; m < len(this.times); m++ {
		if this.times[m] < 0 {
			k--
		} else {
			k++
		}
		res = max(res, k)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
