package p2645minimumadditionstomakevalidstring

func addMinimum(word string) int {
	// Just greedy..
	var res int
	want := rune(0)
	for _, c := range word {
		c -= 'a'
		for c != want {
			res++
			want = (want + 1) % 3
		}
		want = (want + 1) % 3
	}
	res += int(3-want) % 3
	return res
}
