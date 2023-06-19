package p1732findthehighestaltitude

func largestAltitude(gain []int) int {
	var alt int
	var res int
	for _, g := range gain {
		alt += g
		if alt > res {
			res = alt
		}
	}
	return res
}
