package p1491averagesalaryexcludingtheminimumandmaximumsalary

func average(salary []int) float64 {
	min := salary[0]
	max := salary[1]
	var sum int
	for _, s := range salary {
		sum += s
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}
