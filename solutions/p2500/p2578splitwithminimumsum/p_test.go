package p2578splitwithminimumsum

func splitNum(num int) int {
	var count [10]int
	for num > 0 {
		count[num%10]++
		num /= 10
	}
	mul := 1
	var res int
	var i int
	for x := 9; x > 0; x-- {
		for j := 0; j < count[x]; j++ {
			i++
			res += mul * x
			if i%2 == 0 {
				mul *= 10
			}
		}
	}
	return res
}
