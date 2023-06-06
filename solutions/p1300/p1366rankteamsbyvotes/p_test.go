package p1366rankteamsbyvotes

import "sort"

func rankTeams(votes []string) string {
	var count [26][26]int
	for _, v := range votes {
		for i, ch := range v {
			count[ch-'A'][i]++
		}
	}
	cand := []byte(votes[0])
	sort.Slice(cand, func(i, j int) bool {
		a := count[cand[i]-'A']
		b := count[cand[j]-'A']
		for rank := 0; rank < 26; rank++ {
			if a[rank] == b[rank] {
				continue
			}
			return a[rank] > b[rank]
		}
		return cand[i] < cand[j]
	})
	return string(cand)
}
