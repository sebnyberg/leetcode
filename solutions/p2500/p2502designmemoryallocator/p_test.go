package p2502designmemoryallocator

type block struct {
	idx int
	sz  int
}

type Allocator struct {
	mem []int
}

func Constructor(n int) Allocator {
	var a Allocator
	a.mem = make([]int, n)
	return a
}

func (this *Allocator) Allocate(size int, mID int) int {
	var windowSz int
	for i, f := range this.mem {
		if f != 0 {
			windowSz = 0
			continue
		}
		windowSz++
		if windowSz == size {
			for j := i - size + 1; j <= i; j++ {
				this.mem[j] = mID
			}
			return i - size + 1
		}
	}
	return -1
}

func (this *Allocator) Free(mID int) int {
	var count int
	for i := range this.mem {
		if this.mem[i] == mID {
			this.mem[i] = 0
			count++
		}
	}
	return count
}
