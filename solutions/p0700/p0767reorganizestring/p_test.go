package p0767reorganizestring

import "sort"

func reorganizeString(s string) string {
	type charFreq struct {
		ch    rune
		count int32
	}
	freq := make([]charFreq, 26)
	for i := range freq {
		freq[i].ch = rune('a' + i)
	}
	for _, ch := range s {
		freq[ch-'a'].count++
	}

	sortFreq := func() {
		sort.Slice(freq, func(i, j int) bool {
			return freq[i].count > freq[j].count
		})
	}
	res := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		// Add most frequent char to res which is not the previous char
		sortFreq()
		if i == 0 {
			res[i] = freq[0].ch
			freq[0].count--
			continue
		}
		if freq[0].ch != res[i-1] {
			res[i] = freq[0].ch
			freq[0].count--
			continue
		}
		if freq[1].count == 0 {
			return ""
		}
		res[i] = freq[1].ch
		freq[1].count--
	}
	return string(res)
}
