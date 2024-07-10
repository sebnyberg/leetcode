package p1701averagewaitingtime

func averageWaitingTime(customers [][]int) float64 {
	var totalWait int
	var t int
	for _, c := range customers {
		arrival := c[0]
		cookTime := c[1]
		t = max(t+cookTime, arrival+cookTime)
		totalWait += t - arrival
	}
	return float64(totalWait) / float64(len(customers))
}