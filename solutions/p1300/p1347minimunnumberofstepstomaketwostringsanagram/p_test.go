package p1347minimunnumberofstepstomaketwostringsanagram

func minSteps(s string, t string) int {
	var sfreq [26]int
	var tfreq [26]int
	for i := range t {
		sfreq[s[i]-'a']++
		tfreq[t[i]-'a']++
	}
	var delta int
	for i := range sfreq {
		delta += abs(sfreq[i] - tfreq[i])
	}
	return delta / 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
