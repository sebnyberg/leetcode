package p3304findthekthcharacterinastringgamei

func kthCharacter(k int) byte {
	s := []byte{0}
	for len(s) < k {
		n := len(s)
		for i := 0; i < n; i++ {
			s = append(s, (s[i]+1)%26)
		}
	}
	return byte(s[k-1] + 'a')
}
