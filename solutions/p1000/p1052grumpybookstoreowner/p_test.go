package p1052grumpybookstoreowner

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var satisfied int
	for i, c := range customers {
		if grumpy[i] == 0 {
			satisfied += c
		}
	}

	// var grumpy int
	var maxRem int
	var sum int
	for i, c := range customers {
		if grumpy[i] == 1 {
			// grumpy += c
			sum += c
		}
		if i >= minutes {
			if grumpy[i-minutes] == 1 {
				sum -= customers[i-minutes]
			}
		}
		maxRem = max(maxRem, sum)
	}
	return satisfied + maxRem
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
