package p1647minimumdeletionstomakecharacterfrequenciesunique

import "sort"

func minDeletions(s string) int {
	var charFreq [26]int
	for _, ch := range s {
		charFreq[ch-'a']++
	}
	sort.Ints(charFreq[:])
	freqCount := make(map[int]int)
	for _, freq := range charFreq {
		if freq > 0 {
			freqCount[freq]++
		}
	}
	var deletions int
	for freq, c := range freqCount {
		for ; c > 1; c-- {
			var i = freq - 1
			for ; freqCount[i] != 0; i-- {
			}
			if i > 0 {
				freqCount[i] = 1
			}
			deletions += freq - i
		}
	}
	return deletions
}
