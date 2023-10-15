package p2900longestunequaladjacentgroupssubsequencei

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	// Greedy.
	// Just pick any word for which the group does not equal the previously
	// selected group
	prev := -1
	var res []string
	for i, w := range words {
		if groups[i] == prev {
			continue
		}
		res = append(res, w)
		prev = groups[i]
	}
	return res
}
