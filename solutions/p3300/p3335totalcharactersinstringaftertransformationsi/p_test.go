package p3335totalcharactersinstringaftertransformationsi

const mod = 1e9 + 7

func lengthAfterTransformations(s string, t int) int {
	var count [26]int
	for _, c := range s {
		count[c-'a']++
	}
	for range t {
		prev := count[0]
		for j := 1; j < 26; j++ {
			prev, count[j] = count[j], prev
		}
		count[0] = prev
		count[1] = (count[1] + prev) % mod
	}
	var res int
	for _, c := range count {
		res = (res + c) % mod
	}
	return res
}
