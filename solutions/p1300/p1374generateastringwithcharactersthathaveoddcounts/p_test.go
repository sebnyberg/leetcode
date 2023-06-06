package p1374generateastringwithcharactersthathaveoddcounts

func generateTheString(n int) string {
	res := make([]byte, n)
	for i := 0; i < n-(1-n&1); i++ {
		res[i] = 'a'
	}
	for i := n - (1 - n&1); i < n; i++ {
		res[i] = 'b'
	}
	return string(res)
}
