package p1431kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
	res := make([]bool, len(candies))
	for i := range res {
		res[i] = candies[i]+extraCandies >= max
	}
	return res
}
