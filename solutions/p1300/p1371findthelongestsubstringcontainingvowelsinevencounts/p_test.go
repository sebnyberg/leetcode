package p1371findthelongestsubstringcontainingvowelsinevencounts

func findTheLongestSubstring(s string) int {
	first := make(map[int]int)
	var bm int
	first[0] = -1
	bits := [256]int{
		'a': 1,
		'e': 2,
		'i': 3,
		'o': 4,
		'u': 5,
	}
	var res int
	for i, ch := range s {
		if bit := bits[ch]; bit > 0 {
			bm ^= 1 << bits[ch]
		}
		if l, exists := first[bm]; exists {
			res = max(res, i-l)
		} else {
			first[bm] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
