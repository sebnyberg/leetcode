package p2575findthedivisibilityarrayofastring

func divisibilityArray(word string, m int) []int {
	// We can continuously divide by m as numbers are read from word. Only the
	// remainder matters anyway.
	var x int
	n := len(word)
	res := make([]int, n)
	for i := range word {
		x = ((x * 10) + int(word[i]-'0')) % m
		if x == 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}
