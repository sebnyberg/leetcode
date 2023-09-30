package p2739totaldistancetraveled

func distanceTraveled(mainTank int, additionalTank int) int {
	var res int
	for mainTank > 0 {
		res += min(5, mainTank) * 10
		if mainTank >= 5 && additionalTank > 0 {
			mainTank++
			additionalTank--
		}
		mainTank -= 5
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
