package dsx

type RingBuffer struct {
	curLen    int
	maxLen    int
	items     []int
	insertPos int
}

func NewRingBuffer(n int) *RingBuffer {
	return &RingBuffer{
		curLen:    0,
		maxLen:    n,
		items:     make([]int, n),
		insertPos: 0,
	}
}

func (b *RingBuffer) Len() int { return b.curLen }

// Inserts a value into the ring buffer. If an existing element was replaced,
// replaced is set to true, and prev contains the replaced element
func (b *RingBuffer) Insert(val int) (overflow bool, prev int) {
	if b.curLen == b.maxLen {
		overflow = true
		prev = b.items[b.insertPos]
	} else if b.curLen < b.maxLen {
		b.curLen++
	}
	b.items[b.insertPos] = val
	b.insertPos++
	b.insertPos %= b.maxLen
	return overflow, prev
}
