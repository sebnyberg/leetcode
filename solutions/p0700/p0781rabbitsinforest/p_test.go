package p0781rabbitsinforest

func numRabbits(answers []int) int {
	m := make(map[int]int)
	for _, x := range answers {
		m[x+1]++
	}
	var res int
	for x, cnt := range m {
		res += x * (((cnt - 1) / x) + 1)
	}
	return res
}
