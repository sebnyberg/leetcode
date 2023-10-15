package p2904shortestandlexicographicallysmallestbeautifulstring

func shortestBeautifulSubstring(s string, k int) string {
	// Just do greedy. Any valid string must have a rightmost 1, so find the
	// substring that ends with a given 1 + distance to first 1 in the range.
	indices := []int{}
	res := s + "a"
	for i, ch := range s {
		if ch == '0' {
			continue
		}

		indices = append(indices, i)
		if len(indices) > k {
			indices = indices[1:]
		}
		if len(indices) < k {
			continue
		}
		m := indices[k-1] - indices[0] + 1
		ss := s[indices[0] : indices[k-1]+1]
		if m > len(res) || (m == len(res) && ss >= res) {
			continue
		}
		res = ss
	}
	if res == s+"a" {
		return ""
	}
	return res
}
