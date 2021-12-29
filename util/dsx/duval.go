package dsx

import "sort"

// Find the Lyndon words from the alphabet
func duval(alphabet string, maxLen int) []string {
	alphaBytes := []byte(alphabet)
	sort.Slice(alphaBytes, func(i, j int) bool {
		return alphaBytes[i] < alphaBytes[j]
	})
	res := make([]string, 0)

	w := []int{-1}
	for len(w) > 0 {
		// Increment last character
		w[len(w)-1] += 1
		m := len(w)
		if m == maxLen {
			word := make([]byte, m)
			for i := range word {
				word[i] = alphaBytes[w[i]]
			}
			res = append(res, string(word))
		}

		for len(w) < maxLen {
			w = append(w, w[len(w)-m])
		}
		for len(w) > 0 && w[len(w)-1] == len(alphabet)-1 {
			w = w[:len(w)-1]
		}
	}
	return res
}
