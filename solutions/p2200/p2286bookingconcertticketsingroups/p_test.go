package p2286bookingconcertticketsingroups

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		c := Constructor(2, 5)
		a := c.Gather(4, 0)
		require.Equal(t, []int{0, 0}, a)
		a = c.Gather(2, 0)
		require.Equal(t, []int{}, a)
		b := c.Scatter(5, 1)
		require.Equal(t, true, b)
		b = c.Scatter(5, 1)
		require.Equal(t, false, b)
	})
	t.Run("b", func(t *testing.T) {
		c := Constructor(2, 6)
		b := c.Scatter(2, 1)
		require.Equal(t, true, b)
		b = c.Scatter(8, 0)
		require.Equal(t, false, b)
	})
	t.Run("49", func(t *testing.T) {
		c := Constructor(5, 9)
		require.Equal(t, []int{}, c.Gather(10, 1))
		require.Equal(t, true, c.Scatter(3, 3))
		require.Equal(t, []int{1, 0}, c.Gather(9, 1))
		require.Equal(t, []int{}, c.Gather(10, 2))
		require.Equal(t, []int{0, 3}, c.Gather(2, 0))
	})
}

type BookMyShow struct {
	rows           int
	seats          int
	m              int
	tree           []*treeNode
	firstWithSeats int
}

type treeNode struct {
	seats   int
	seatSum int
}

func Constructor(m int, n int) BookMyShow {
	// NOTE!!!! I swapped m and n in the names of the input because Leetcode is
	// not conforming to standards. N is always the number of elements per row,
	// and m is the number of rows. Jeez.

	// This solution uses a segment tree to keep track of maximum number of
	// available seats per row for a given range of rows.

	var s BookMyShow
	s.rows = m
	s.seats = n

	// From now on, m == treeSize
	m = 1
	for m < s.rows {
		m <<= 1
	}
	s.m = m
	s.tree = make([]*treeNode, m*2)
	for i := s.m; i < s.m*2; i++ {
		s.tree[i] = &treeNode{
			seats:   s.seats,
			seatSum: s.seats,
		}
	}
	for i := s.m - 1; i >= 1; i-- {
		s.tree[i] = &treeNode{
			seats:   s.seats,
			seatSum: s.tree[i*2].seatSum + s.tree[i*2+1].seatSum,
		}
	}
	s.firstWithSeats = 0
	return s
}

func (this *BookMyShow) update(i int, seats int) {
	this.tree[i].seats = seats
	this.tree[i].seatSum = seats
	for j := i / 2; j >= 1; j /= 2 {
		this.tree[j].seats = max(this.tree[j*2].seats, this.tree[j*2+1].seats)
		this.tree[j].seatSum = this.tree[j*2].seatSum + this.tree[j*2+1].seatSum
	}
}

// search returns the lowest index for which the value is val
// in this case this will find the row that contains >= the requested number of
// seats.
func (this *BookMyShow) search(i int, lo, hi, qlo, qhi, val int) int {
	if this.tree[i].seats < val {
		return -1
	}
	if lo == hi {
		// At a leaf node
		return lo
	}
	mid := lo + (hi-lo)/2
	if res := this.search(i*2, lo, mid, qlo, qhi, val); res != -1 {
		return res
	}
	if mid+1 > qhi {
		return -1
	}
	return this.search(i*2+1, mid+1, hi, mid, qhi, val)
}

// querySum return the sum of the range [qlo,qhi]
func (this *BookMyShow) querySum(i int, lo, hi, qlo, qhi int) int {
	if qhi < lo || qlo > hi {
		return 0
	}
	if qlo <= lo && qhi >= hi {
		return this.tree[i].seatSum
	}
	mid := lo + (hi-lo)/2
	left := this.querySum(i*2, lo, mid, qlo, qhi)
	right := this.querySum(i*2+1, mid+1, hi, qlo, qhi)
	return left + right
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	// Search for the first row that fits k people.
	row := this.search(1, 0, this.m-1, 0, maxRow, k)
	if row == -1 {
		return []int{}
	}
	seatsFree := this.tree[this.m+row].seats
	nleft := seatsFree - k
	this.update(this.m+row, nleft)
	return []int{row, this.seats - seatsFree}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	seatsAvailable := this.querySum(1, 0, this.m-1, 0, maxRow)
	if seatsAvailable < k {
		return false
	}
	for {
		d := min(k, this.tree[this.m+this.firstWithSeats].seats)
		k -= d
		x := this.tree[this.m+this.firstWithSeats].seats
		this.update(this.m+this.firstWithSeats, x-d)
		if k == 0 {
			break
		}
		this.firstWithSeats++
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
