package p0957prisoncellsafterndays

func prisonAfterNDays(cells []int, n int) []int {
	m := len(cells)
	var state [8]byte
	for i := range cells {
		state[i] = byte(cells[i])
	}
	seen := make(map[[8]byte]int)
	states := make([][8]byte, 0)
	states = append(states, state)
	seen[state] = 0
	var repeatIdx int
	for {
		var next [8]byte
		for i := 1; i < m-1; i++ {
			next[i] = ^(state[i-1] ^ state[i+1]&1) & 1
		}
		if i, exists := seen[next]; exists {
			repeatIdx = i
			break
		}
		seen[next] = len(states)
		states = append(states, next)
		state = next
	}
	asresult := func(s [8]byte) []int {
		var res []int
		for _, v := range s[:m] {
			res = append(res, int(v))
		}
		return res
	}
	if n >= len(states) {
		n = repeatIdx + (n-repeatIdx)%(len(states)-repeatIdx)
	}
	return asresult(states[n])
}
