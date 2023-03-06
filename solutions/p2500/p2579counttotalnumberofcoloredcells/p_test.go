package p2579counttotalnumberofcoloredcells

func coloredCells(n int) int64 {
	if n == 1 {
		return 1
	}
	res := 1
	extra := 0
	for i := 2; i <= n; i++ {
		res += 4 + extra
		extra += 4
	}
	return int64(res)
}
