For each number in `nums2`, find the index that number would have if it were in
`nums1`. Let's call it n1Pos:

```go
nums1 = [ 0, 5, 1, 4, 2, 3 ]
nums2 = [ 5, 1, 0, 4, 3, 2 ]
n1Pos = [ 1, 2, 0, 3, 5, 4 ]
```

Consider the index `j=1` in `n1Pos` - let's call it the pivot. 
Any pair `(i, j)` where `i < j` is valid IFF `n1Pos[i] < n1Pos[j]`.
Any pair `(j, k)` where `j < k` is valid IFF `n1Pos[j] < n1Pos[k]`.

Any valid pair on the left can be combined with a valid pair on the right to
form a valid triplet.

The tricky part is that even calculating all possible 2-pairs is too slow, so we
need some faster way of finding the number of values below / above a certain value.

By having a Fenwick tree for the left side, and one for the right side, we can
quickly calculate how many indices were above / below the pivot.

# Solution

```go
func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	idx := make([]int, n)
	for i := range nums1 {
		idx[nums1[i]] = i
	}
	n1Pos := make([]int, n)
	for i, n2 := range nums2 {
		n1Pos[i] = idx[n2]
	}
	var count int64
	right := NewFenwick(n)
	left := NewFenwick(n)
	for i := 2; i < n; i++ {
		right.Add(n1Pos[i], 1)
	}
	left.Add(n1Pos[0], 1)

	for i := 1; i < n-1; i++ {
		below := left.Sum(n1Pos[i])
		above := right.Sum(n-1) - right.Sum(n1Pos[i])
		count += int64(below * above)
		left.Add(n1Pos[i], 1)
		right.Add(n1Pos[i+1], -1)
	}

	return count
}

type Fenwick struct {
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n+1),
	}
}

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

func (f *Fenwick) Add(i int, k int) {
	i++
	for i < len(f.tree) {
		f.tree[i] += k
		i += i & -i
	}
}

func (f *Fenwick) Sum(i int) int {
	i++
	res := 0
	for i > 0 {
		res += f.tree[i]
		i -= i & -i
	}
	return res

}
```