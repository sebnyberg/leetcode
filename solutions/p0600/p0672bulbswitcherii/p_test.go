package p0672bulbswitcherii

func flipLights(n int, presses int) int {
	switch {
	case presses == 0:
		return 1
	case n == 1:
		return 2
	case n == 2 && presses == 1:
		return 3
	case n == 2, presses == 1:
		return 4
	case presses == 2:
		return 7
	default:
		return 8
	}
}
