package p2938separateblackandwhiteballs

func minimumSteps(s string) int64 {
	//
	// 0001101101
	//  0001011011
	//   0000110111
	//
	// It seems like switching a given 1 into its rightmost location costs the
	// number of zeroes to the right of its current location.
	//
	n := len(s)
	right := make([]int, n+1)
	for i := len(s) - 1; i >= 0; i-- {
		right[i] = right[i+1]
		if s[i] == '0' {
			right[i]++
		}
	}
	var cost int
	for i := range s {
		if s[i] == '1' {
			cost += right[i]
		}
	}
	return int64(cost)
}
