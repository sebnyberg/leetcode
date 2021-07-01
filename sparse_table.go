package leetcode

// SparseTable is a construct which allows for range min/max queries (RMQ) in
// O(1). Construction of the table is done in O(n*logn). The basic idea of the
// is to construct sets of intervals which are powers of two in width. Starting
// with width 1, the first level is simply the input numbers. At level 2, the
// minimum value of width-1 intervals is put into length 2 intervals, and so on,
// until the next interval width would exceed the length of the input array.
type SparseTable struct {
	table [][]int
	cmp   func(a, b int) int
}

func NewSparseTable(nums []int, cmp func(a, b int) int) *SparseTable {
	n := len(nums)
	k := intlog2(n)
	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, k+1)
	}

	// Initialize width-1
	for i, n := range nums {
		table[i][0] = n
	}

	// Calculate other intervals from 2^1, 2^2, ..., 2^k
	for i := 1; i <= k; i++ {
		for j := 0; j+(1<<i)-1 < n; j++ {
			table[j][i] = cmp(
				table[j][i-1],
				table[j+(1<<(i-1))][i-1],
			)
		}
	}

	return &SparseTable{
		table: table,
		cmp:   cmp,
	}
}

// Query returns the min or max value within the provided range [start, end].
// Whether min or max is returned depends on the cmp function that was provided
// to NewSparseTable(...).
func (t *SparseTable) Query(start, end int) int {
	start = max(start, 0)
	end = min(end, len(t.table)-1)
	width := end - start + 1
	k := intlog2(width)
	return t.cmp(t.table[start][k], t.table[end-(1<<k)][k])
}

// intlog2 returns the greatest number 'a' such that 2^a < x.
func intlog2(x int) int {
	a := 0
	v := 2
	for v < x {
		v <<= 1
		a++
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
