package p1103distributecandiestopeople

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)

	x := 1
	var i int
	for candies > 0 {
		y := min(x, candies)
		res[i] += y
		candies -= y
		x++
		i = (i + 1) % num_people
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
