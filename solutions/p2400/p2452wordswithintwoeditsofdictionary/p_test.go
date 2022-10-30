package p2452wordswithintwoeditsofdictionary

func twoEditWords(queries []string, dictionary []string) []string {
	comp := func(w1, w2 string) bool {
		var diff int
		for i := range w1 {
			if w1[i] != w2[i] {
				diff++
				if diff > 2 {
					return false
				}
			}
		}
		return true
	}
	var res []string
	for _, w := range queries {
		for _, w2 := range dictionary {
			if comp(w, w2) {
				res = append(res, w)
				break
			}
		}
	}
	return res
}
