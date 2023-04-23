package p2582passthepillow

func passThePillow(n int, time int) int {
	time %= ((n - 1) * 2)
	if time >= n-1 {
		return n - (time - (n - 1))
	}
	return time + 1
}
