package p0858mirrorreflection

func mirrorReflection(p int, q int) int {
	var bounces byte
	var dy int
	for {
		bounces++
		dy = (dy + q) % (2 * p)
		leftSide := bounces&1 == 0
		if leftSide {
			if dy == p {
				return 2
			}
		} else {
			if dy == 0 {
				return 0
			}
			if dy == p {
				return 1
			}
		}
	}
}
