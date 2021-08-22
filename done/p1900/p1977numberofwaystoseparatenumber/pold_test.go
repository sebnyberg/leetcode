package p1977numberofwaystoseparatenumber_test

const mod = 1e9 + 7

// dp[end] = make([]lastNumCount, 0, len(dp[end-1])/2)
// presums[end] = make([]int, 0, len(presums[end-1])/2)
func numberOfCombinations(num string) int {
	type numSeries struct {
		val   string
		count int
		pos   int
	}
	series := []numSeries{{"0", 1, 0}}
	for i := 1; i <= len(num); i++ {
		newSeries := map[string]int{}
		for _, s := range series {
			if geq(num[s.pos:i], s.val) && num[s.pos] != '0' {
				newSeries[num[s.pos:i]] += s.count
				newSeries[num[s.pos:i]] %= mod
			}
		}
		for val, count := range newSeries {
			series = append(series, numSeries{val, count, i})
		}
	}
	var res int
	for _, s := range series {
		if s.pos == len(num) {
			res += s.count
			res %= mod
		}
	}
	return res % mod
}

func geq(a, b string) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] > b[i]
	}
	return true
}
