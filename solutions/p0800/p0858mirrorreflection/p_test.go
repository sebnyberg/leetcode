package p0858mirrorreflection

func mirrorReflection(p int, q int) int {
	// Due to the pidgeon-hole principle, with the constraints,
	// we can do a forever-loop and still be sure that the loop
	// will finish eventually
	var bounces byte
	var d int
	for {
		bounces++
		d = (d + q) % (2 * p)
		if bounces&1 == 0 {
			// Left side
			if d == p {
				return 2
			}
		} else {
			// Right side
			if d == 0 {
				return 0
			}
			if d == p {
				return 1
			}
		}
	}
}
