package p2438rangeproductqueriesofpowers

func productQueries(n int, queries [][]int) []int {
	const mod = 1e9 + 7
	powers := []int{}
	for factor := 1; factor <= n; factor <<= 1 {
		if n&factor > 0 {
			powers = append(powers, factor)
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		v := powers[q[0]]
		for i := q[0] + 1; i <= q[1]; i++ {
			v = (v * powers[i]) % mod
		}
		res[i] = v
	}
	return res
}
