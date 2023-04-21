package p1208getequalsubstringswithinbudget

func equalSubstring(s string, t string, maxCost int) int {
	var res int
	var j int
	var cost int
	for i := range s {
		cost += abs(int(s[i]) - int(t[i]))
		for cost > maxCost {
			cost -= abs(int(s[j]) - int(t[j]))
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
