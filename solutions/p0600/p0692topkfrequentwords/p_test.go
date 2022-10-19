package p0692topkfrequentwords

import "sort"

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	type wordCount struct {
		word  string
		count int
	}
	wcs := make([]wordCount, 0, len(m))
	for w, count := range m {
		wcs = append(wcs, wordCount{w, count})
	}
	sort.Slice(wcs, func(i, j int) bool {
		if wcs[i].count == wcs[j].count {
			return wcs[i].word < wcs[j].word
		}
		return wcs[i].count > wcs[j].count
	})
	res := make([]string, min(k, len(wcs)))
	for i := 0; i < min(len(wcs), k); i++ {
		res[i] = wcs[i].word
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
