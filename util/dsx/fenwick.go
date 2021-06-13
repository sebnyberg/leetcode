package ds

type Fenwick struct {
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n),
	}
}

// Init initializes the Fenwick tree, overwriting old content with
// the provided list
func (f *Fenwick) Init(vals []int) {
	n := len(vals)
	copy(f.tree, vals)
	for i := 1; i < n; i++ {
		p := i + (i & -i)
		if p < n {
			f.tree[p] += f.tree[i]
		}
	}
}

// Add k to index i
func (f *Fenwick) Add(i int, k int) {
	for i < len(f.tree) {
		f.tree[i] += k
		i += i & -i
	}
}

func (f *Fenwick) Sum(i int) int {
	res := 0
	for i > 0 {
		res += f.tree[i]
		// flip last set bit to zero (use parent in tree)
		i -= i & -i
	}
	return res
}

func (f *Fenwick) SumRange(from int, to int) int {
	return f.Sum(to) - f.Sum(from)
}
