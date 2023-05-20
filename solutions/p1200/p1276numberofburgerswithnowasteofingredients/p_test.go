package p1276numberofburgerswithnowasteofingredients

func numOfBurgers(tomato int, cheese int) []int {
	if tomato%2 == 1 || tomato > cheese*4 || tomato < cheese*2 {
		return []int{}
	}
	// We will create 'cheese' number of burgers
	// If we made only small burgers, we would use 2*cheese tomatoes
	// Each large burger increases total tomato consumption by 2
	// Therefore, the number of large burgers is equal to number of extra
	// tomatoes from making small burgers, divided by two.
	tomato -= cheese * 2
	nlarge := tomato / 2
	nsmall := cheese - nlarge
	return []int{nlarge, nsmall}
}
