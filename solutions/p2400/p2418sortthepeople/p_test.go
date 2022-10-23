package p2418sortthepeople

import "sort"

type peopleSorter struct {
	names   []string
	heights []int
}

func (p *peopleSorter) Swap(i, j int) {
	p.heights[i], p.heights[j] = p.heights[j], p.heights[i]
	p.names[i], p.names[j] = p.names[j], p.names[i]
}

func (p peopleSorter) Len() int {
	return len(p.names)
}
func (p peopleSorter) Less(i, j int) bool {
	return p.heights[i] > p.heights[j]
}

func sortPeople(names []string, heights []int) []string {
	s := &peopleSorter{
		names:   names,
		heights: heights,
	}
	sort.Sort(s)
	return names
}
