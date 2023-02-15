package p1146snapshotarray

import (
	"sort"
	"testing"
)

func TestSnapshotArray(t *testing.T) {
	a := Constructor(3)
	a.Set(0, 5)
	a.Snap()
	a.Set(0, 6)
	res := a.Get(0, 0)
	_ = res
}

type SnapshotArray struct {
	vals    [][]int
	rev     [][]int
	version int
}

func Constructor(length int) SnapshotArray {
	var sa SnapshotArray
	sa.vals = make([][]int, length)
	sa.rev = make([][]int, length)
	for i := range sa.vals {
		sa.vals[i] = append(sa.vals[i], 0)
		sa.rev[i] = append(sa.rev[i], 0)
	}
	return sa
}

func (this *SnapshotArray) Set(index int, val int) {
	i := index
	j := len(this.vals[index]) - 1
	if this.rev[i][j] < this.version {
		this.rev[i] = append(this.rev[i], this.version)
		this.vals[i] = append(this.vals[i], val)
	} else {
		this.vals[i][j] = val
	}
}

func (this *SnapshotArray) Snap() int {
	this.version++
	return this.version - 1
}

func (this *SnapshotArray) Get(index int, snapID int) int {
	j := sort.SearchInts(this.rev[index], snapID+1) - 1
	return this.vals[index][j]
}
