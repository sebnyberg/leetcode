package p2416sumofprefixscoresofstrings

type trieNode struct {
	count int
	next  [26]*trieNode
}

func sumPrefixScores(words []string) []int {
	root := &trieNode{}
	for _, w := range words {
		curr := root
		for _, ch := range w {
			c := int(ch - 'a')
			if curr.next[c] == nil {
				curr.next[c] = &trieNode{}
			}
			curr = curr.next[c]
			curr.count++
		}
	}
	n := len(words)
	res := make([]int, n)
	for i, w := range words {
		curr := root
		for _, ch := range w {
			curr = curr.next[ch-'a']
			res[i] += curr.count
		}
	}
	return res
}
