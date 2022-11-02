package p0893groupsofspecialequivalentstrings

func numSpecialEquivGroups(words []string) int {
	// This sounds harder than it is
	// Two strings are equivalent if their character "footprint" is the same for
	// even and odd indices alike.
	//
	footprint := func(w string) [2][26]int {
		var res [2][26]int
		for i := range w {
			res[i&1][w[i]-'a']++
		}
		return res
	}

	m := make(map[[2][26]int][]string)
	for _, w := range words {
		m[footprint(w)] = append(m[footprint(w)], w)
	}
	return len(m)
}
