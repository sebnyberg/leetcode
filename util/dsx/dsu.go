package dsx

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) Find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.Find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) Union(a, b int) {
	a = d.Find(a)
	b = d.Find(b)
	if a != b {
		if d.size[a] < d.size[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
	}
}
