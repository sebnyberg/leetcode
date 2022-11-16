package p0984stringrithoutaaaorbbb

func strWithout3a3b(a int, b int) string {
	n := a + b
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		if a > b {
			if i <= 1 || res[i-1] != 'a' || res[i-2] != 'a' {
				res[i] = 'a'
				a--
			} else {
				res[i] = 'b'
				b--
			}
		} else {
			if i <= 1 || res[i-1] != 'b' || res[i-2] != 'b' {
				res[i] = 'b'
				b--
			} else {
				res[i] = 'a'
				a--
			}
		}
	}

	return string(res)
}
