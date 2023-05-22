package p1397findallgoodstrings

const mod = 1e9 + 7

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	lps := buildLPS(evil)
	mem := make(map[key]int)
	res := dfs(mem, s1, s2, evil, lps, true, true, 0, 0, n)
	return res
}

func dfs(mem map[key]int, s1, s2, evil string, lps []int, lowerBound, upperBound bool, i, j, n int) int {
	if i == n {
		return 1
	}
	k := key{
		i:          i,
		j:          j,
		lowerBound: lowerBound,
		upperBound: upperBound,
	}
	if v, exists := mem[k]; exists {
		return v
	}
	var lo byte = 'a'
	if lowerBound {
		lo = s1[i]
	}
	var hi byte = 'z'
	if upperBound {
		hi = s2[i]
	}
	var res int
	for ch := lo; ch <= hi; ch++ {
		k := j
		for k > 0 && evil[k] != ch {
			k = lps[k-1]
		}
		if ch == evil[k] {
			k++
			if k == len(evil) {
				continue
			}
		}
		lb := lowerBound && ch == s1[i]
		ub := upperBound && ch == s2[i]
		res += dfs(mem, s1, s2, evil, lps, lb, ub, i+1, k, n)
		res %= mod
	}

	mem[k] = res
	return res
}

type key struct {
	i          int
	j          int
	lowerBound bool
	upperBound bool
}

func buildLPS(s string) []int {
	lps := make([]int, len(s))
	lps[0] = 0
	i, j := 0, 1
	for j < len(s) {
		if s[i] == s[j] {
			i++
			lps[j] = i
			j++
		} else if i == 0 {
			lps[j] = 0
			j++
		} else {
			i = lps[i-1]
		}
	}
	return lps
}
