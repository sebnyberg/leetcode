package p2456mostpopularvideocreator

import "sort"

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	viewsTotal := make(map[string]int)
	mostViewed := make(map[string]int)
	for i := range creators {
		c := creators[i]
		v := views[i]
		viewsTotal[c] += v
		if vi, exists := mostViewed[c]; !exists || views[vi] < v || views[vi] == v && ids[i] < ids[vi] {
			mostViewed[c] = i
		}
	}
	type result struct {
		name  string
		id    string
		count int
	}
	results := make([]result, 0, len(viewsTotal))
	for name := range viewsTotal {
		results = append(results, result{
			name:  name,
			id:    ids[mostViewed[name]],
			count: viewsTotal[name],
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})
	ret := [][]string{
		{results[0].name, results[0].id},
	}
	for i := 1; i < len(results) && results[i].count == results[i-1].count; i++ {
		ret = append(ret, []string{results[i].name, results[i].id})
	}
	return ret
}
